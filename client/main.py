#!python3
import sys
import requests
import csv
import os


DATA_FILE = "db.csv"
METADATA_FILE = "last_sync.conf"


def init_db_if_not_exists():
    if not os.path.exists(DATA_FILE):
        with open(DATA_FILE, "w", encoding="utf8", newline="") as f:
            csvWriter = csv.DictWriter(
                f, fieldnames=["uuid", "author", "message", "likes"]
            )
            csvWriter.writeheader()

    if not os.path.exists(METADATA_FILE):
        with open(METADATA_FILE, "w") as f:
            f.write("0")


def fetch_update(url: str) -> dict:
    with open(METADATA_FILE, "r") as f:
        last_sync_timestamp = int(f.read())

    r = requests.get(url, params={"last_sync_timestamp": last_sync_timestamp})
    if r.status_code != 200:
        print("Error: %s" % r.status_code)
        sys.exit(1)
    last_sync_timestamp = r.headers["Last-Sync"]

    with open(METADATA_FILE, "w") as f:
        f.write(str(last_sync_timestamp))
    return r.json()


def read_records() -> dict[str, dict[str, str]]:
    records = {}
    with open(DATA_FILE, "r", encoding="utf8") as f:
        csvReader = csv.DictReader(f)
        for row in csvReader:
            records[row["uuid"]] = row
    return records


def sync_records(records: dict[str, dict[str, str]], updates: list[dict]):
    for update in updates:
        uuid = update["uuid"]
        if uuid in records:
            records[uuid].update(update)
        else:
            records[uuid] = update


def write_records(records: dict[str, dict[str, str]]) -> None:
    with open(DATA_FILE, "w", encoding="utf8", newline="") as f:
        csvWriter = csv.DictWriter(f, fieldnames=["uuid", "author", "message", "likes"])
        csvWriter.writeheader()
        for record in records.values():
            csvWriter.writerow(record)


def main():
    if len(sys.argv) != 2:
        print("Usage: ./main.py <URL>")
        sys.exit(1)
    url = sys.argv[1]

    init_db_if_not_exists()

    records = read_records()
    updates = fetch_update(url)
    sync_records(records, updates)
    write_records(records)


if __name__ == "__main__":
    main()
