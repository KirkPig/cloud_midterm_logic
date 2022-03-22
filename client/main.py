import json
import math
import sys
import threading
import requests
import csv
import os
import fastavro
from io import BytesIO
from datetime import datetime


DATA_FILE = "state.csv"
METADATA_FILE = "last_sync.conf"
SYNC_RECORD_LIMIT = 60000

schema = fastavro.parse_schema(json.load(open("avro/sync_message.avsc", "r")))

def deserialize_avro(content: bytes):
    record = fastavro.schemaless_reader(BytesIO(content), schema)
    return record


def init_db_if_not_exists():
    if not os.path.exists(DATA_FILE):
        with open(DATA_FILE, "w", encoding="utf8", newline="") as f:
            csvWriter = csv.DictWriter(
                f, fieldnames=["uuid", "author", "message", "likes"]
            )
            csvWriter.writeheader()
            print(datetime.now().isoformat(), "Initialized", DATA_FILE)

    if not os.path.exists(METADATA_FILE):
        with open(METADATA_FILE, "w") as f:
            f.write("0")
            print(datetime.now().isoformat(), "Initialized", METADATA_FILE)

def fetch_message_count(base_url: str, last_sync: int) -> int:
    r = requests.get(f"{base_url}/count", params={
        "timestamp": last_sync,
    })
    return int(r.text)

def fetch_update(base_url: str, last_sync: int, offset: int, threadNo: int, threadUpdates: list, threadLastSyncs: list):
    print(datetime.now().isoformat(), f"Fetching {offset} to {offset + SYNC_RECORD_LIMIT}")
    start = datetime.now()
    r = requests.get(f"{base_url}/messages", params={
        "timestamp": last_sync,
        "limit": SYNC_RECORD_LIMIT,
        "offset": offset,
    })
    stop = datetime.now()
    if r.status_code != 200:
        print(datetime.now().isoformat(), "Error: %s" % r.status_code)
        sys.exit(1)

    # record = deserialize_avro(r.content)
    # threadUpdates[threadNo] = record["updates"]
    threadUpdates[threadNo] = r.json()
    print(datetime.now().isoformat(), f"Fetched {offset} to {offset + SYNC_RECORD_LIMIT}. Took {(stop - start).total_seconds()} seconds. Got {len(threadUpdates[threadNo])} updates")
    threadLastSyncs[threadNo] = r.headers["Last-Sync"]
    return


def read_records() -> dict:
    records = {}
    with open(DATA_FILE, "r", encoding="utf8") as f:
        csvReader = csv.DictReader(f)
        for row in csvReader:
            records[row["uuid"]] = row
    print(datetime.now().isoformat(), f"Read {len(records)} records from {DATA_FILE}")
    return records

def sync_records(records: dict, threadUpdates: list):
    for updates in threadUpdates:
        for update in updates:
            uuid = update["uuid"]
            if "isDeleted" in update and uuid in records:
                del records[uuid]
            elif uuid in records:
                if "isDeleted" in update:
                    del update["isDeleted"]
                records[uuid].update(update)
            else:
                if "isDeleted" in update:
                    del update["isDeleted"]
                records[uuid] = update
        print(datetime.now().isoformat(), f"Applied {len(updates)} updates. Current record count: {len(records)}")
    print(datetime.now().isoformat(), f"Applied updates to {len(records)} records")


def write_records(records: dict) -> None:
    with open(DATA_FILE, "w", encoding="utf8", newline="") as f:
        csvWriter = csv.DictWriter(f, fieldnames=["uuid", "author", "message", "likes"])
        csvWriter.writeheader()
        for record in records.values():
            csvWriter.writerow(record)
        print(datetime.now().isoformat(), f"Wrote {len(records)} records to {DATA_FILE}")


def main():
    if len(sys.argv) != 2:
        print("Usage: ./main.py <BASE_URL>")
        print("Example: ./main.py http://54.218.59.213/api")
        sys.exit(1)
    base_url = sys.argv[1]

    init_db_if_not_exists()
    records = read_records()

    with open(METADATA_FILE, "r") as f:
        last_sync = int(f.read())
    count = fetch_message_count(base_url, last_sync)
    if count == 0:
        return

    noThreads = math.ceil(count / SYNC_RECORD_LIMIT)
    threadUpdates = [None] * noThreads
    threadLastSyncs = [None] * noThreads
    threads: list[threading.Thread] = []
    print(f"Starting {noThreads} threads")
    for threadNo in range(noThreads):
        t = threading.Thread(target=fetch_update, args=(base_url, last_sync, threadNo * SYNC_RECORD_LIMIT, threadNo, threadUpdates, threadLastSyncs))
        t.start()
        threads.append(t)

    for threadNo in range(noThreads):
        threads[threadNo].join()

    with open(METADATA_FILE, "w") as f:
        f.write(str(threadLastSyncs[-1]))

    sync_records(records, threadUpdates)
    write_records(records)


if __name__ == "__main__":
    main()
