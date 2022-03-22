#!python3
import csv
import requests
from urllib.parse import urljoin

baseUrl = ""
filename = ""


def deleteMessage(data):
    # print(data['action'], urljoin(baseUrl, urljoin('messages/', data['uuid'])))
    return requests.delete(urljoin(baseUrl, urljoin("messages/", data["uuid"])))


def updateMessage(data):
    # print(urljoin(baseUrl, urljoin(
    #     'messages/', data['uuid'])))
    dataToSend = {
        "author": data["author"],
        "message": data["message"],
        "likes": int(data["likes"]),
    }

    return requests.put(
        urljoin(baseUrl, urljoin("messages/", data["uuid"])), json=dataToSend
    )
    # return (r.status_code == 204, r.text)


def createMessage(data):
    dataToSend = {
        "uuid": data["uuid"],
        "author": data["author"],
        "message": data["message"],
        "likes": int(data["likes"]),
    }

    return requests.post(urljoin(baseUrl, "messages"), json=dataToSend)
    # return (r.status_code == 201, r.text)


def main():
    csvFile = open(filename)
    reader = csv.DictReader(csvFile)
    rowCounter = 0
    for row in reader:
        retryCounter = 0
        while True:
            if retryCounter == 10:
                raise Exception("Message failed to send")

            if row["action"] == "create":
                response = createMessage(row)
                if response.status_code == 201 or response.status_code == 409:
                    print("Row", rowCounter, "sent!")
                    break
                else:
                    print(
                        "Row",
                        rowCounter,
                        "failed to send",
                        response.status_code,
                        ":",
                        response.text,
                    )

            if row["action"] == "update":
                response = updateMessage(row)
                if response.status_code == 204:
                    print("Row", rowCounter, "sent!")
                    break
                elif response.status_code == 404:
                    raise Exception("UUID not found")
                else:
                    print(
                        "Row",
                        rowCounter,
                        "failed to send",
                        response.status_code,
                        ":",
                        response.text,
                    )

            if row["action"] == "delete":
                response = deleteMessage(row)
                if response.status_code == 204 or response.status_code == 404:
                    print("Row", rowCounter, "sent!")
                    break
                else:
                    print(
                        "Row",
                        rowCounter,
                        "failed to send",
                        response.status_code,
                        ":",
                        response.text,
                    )

            retryCounter += 1
        rowCounter += 1

    csvFile.close()


if __name__ == "__main__":
    filename = input("Filename : ")
    baseUrl = input("Base URL with trailing / (e.g. http://10.2.110.61/api/) : ")
    main()
