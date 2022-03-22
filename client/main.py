import json
import sys
import requests
import csv
import os
import fastavro
from io import BytesIO
from datetime import datetime


DATA_FILE = "state.csv"
METADATA_FILE = "last_sync.conf"
SYNC_RECORD_LIMIT = 600000

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

def fetch_update(base_url: str) -> dict:
    with open(METADATA_FILE, "r") as f:
        last_sync = int(f.read())

    count = fetch_message_count(base_url, last_sync)
    print(datetime.now().isoformat(), "Total update count:", count)

    updates = []
    offset = 0
    new_last_sync = last_sync
    while offset < count:
        print(datetime.now().isoformat(), f"Fetching {offset} to {offset + SYNC_RECORD_LIMIT}")
        start = datetime.now()
        r = requests.get(f"{base_url}/messages", params={
            "timestamp": last_sync,
            "limit": SYNC_RECORD_LIMIT,
            "offset": offset,
        })
        stop = datetime.now()
        print(datetime.now().isoformat(), f"Fetched {offset} to {offset + SYNC_RECORD_LIMIT}. Took {(stop - start).total_seconds()} seconds")
        if r.status_code != 200:
            print(datetime.now().isoformat(), "Error: %s" % r.status_code)
            sys.exit(1)
        record = deserialize_avro(r.content)
        updates.extend(record["updates"])
        
        offset += SYNC_RECORD_LIMIT

        new_last_sync = r.headers["Last-Sync"]
    last_sync = new_last_sync

    with open(METADATA_FILE, "w") as f:
        f.write(str(last_sync))
        
    return updates


def read_records() -> dict:
    records = {}
    with open(DATA_FILE, "r", encoding="utf8") as f:
        csvReader = csv.DictReader(f)
        for row in csvReader:
            records[row["uuid"]] = row
    print(datetime.now().isoformat(), f"Read {len(records)} records from {DATA_FILE}")
    return records


def sync_records(records: dict, updates: list):
    for update in updates:
        uuid = update["uuid"]
        if "isDeleted" in update and update["isDeleted"] and uuid in records:
            del records[uuid]
        elif uuid in records:
            del update["isDeleted"]
            records[uuid].update(update)
        else:
            del update["isDeleted"]
            records[uuid] = update
    print(datetime.now().isoformat(), f"Applied {len(updates)} updates to {len(records)} records")


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
    updates = fetch_update(base_url)
    sync_records(records, updates)
    write_records(records)


if __name__ == "__main__":
    main()
