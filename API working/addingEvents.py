import psycopg2
import datetime
import requests
from config import host, user, password, db_name
def insertIntoDatabase(in_events_id, in_event_name, in_event_time, in_event_result, in_subs_id):
    try:
        connection = psycopg2.connect(
            host=host,
            user = user,
            password = password
        )
        
        with connection.cursor() as cursor:
            in_event_time = datetime.datetime.fromisoformat(in_event_time)
            sql1 = "INSERT INTO events (events_id, event_name, event_time, event_result) VALUES (%s, %s, %s, %s)"
            cursor.execute(sql1, (in_events_id, in_event_name, in_event_time, in_event_result))
            sql2 = "INSERT INTO event_sub(event_id, subs_id) VALUES (%s, %s)"
            cursor.execute(sql2, (in_events_id, in_subs_id))
            # insert into events and dont forget to bind event_id with subscription_id
            #sql1 = "INSERT INTO events (event_name, event_time, event_result) VALUES (%s, %s, %s)"
            #cursor.execute(sql1, (in_event_name, in_event_time, in_event_result))
            
        connection.commit()
            
    except Exception as _ex:
        print("INFO  error", _ex)
    finally:
        if connection:
            connection.close()

pageNum = 1
while(True):
    url ="https://api.allsportdb.com/v3/calendar?page=" + str(pageNum) 
    response = requests.get(url, headers={'Authorization': 'Bearer 88d3ca9b-3004-484f-9669-a9112c115637'})
    jsonF = response.json()
    print(jsonF)
    if(jsonF == [] or pageNum == 15):
        break
    for i in jsonF:
        insertIntoDatabase(i["id"], i["name"], i["dateFrom"], "", i["sportId"])
        print(i["id"], i["name"], i["dateFrom"], "", i["sportId"], end="\n\n")
    pageNum += 1