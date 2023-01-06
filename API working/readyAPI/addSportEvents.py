import psycopg2
import datetime
import requests
from config import host, user, password, db_name
def insertIntoDatabase(in_event_name, in_event_timestart, in_event_timeend, in_event_result, in_subs_name):
    try:
        connection = psycopg2.connect(
            host=host,
            user = user,
            password = password
        )
        
        with connection.cursor() as cursor:
            in_event_timestart = datetime.datetime.fromisoformat(in_event_timestart)
            in_event_timeend = datetime.datetime.fromisoformat(in_event_timeend)
            sql1 = "INSERT INTO events (event_name,event_timestart, event_timeend,event_result) VALUES (%s, %s, %s, %s) returning events_id"
            cursor.execute(sql1, (in_event_name, in_event_timestart, in_event_timeend, in_event_result))
            event_id = cursor.fetchone()[0]
            print(in_subs_name)

            
            sql2 = "select subs_id from subscriptions where subs_name = (%s)"
            cursor.execute(sql2, (in_subs_name,))
            subs_id = cursor.fetchone()[0]
            
            sql3 = "INSERT INTO event_sub(event_id, subs_id) VALUES (%s, %s)"
            cursor.execute(sql3, (event_id, subs_id))
            # insert into events and dont forget to bind event_id with subscription_id
            #sql1 = "INSERT INTO events (event_name, event_time, event_result) VALUES (%s, %s, %s)"
            #cursor.execute(sql1, (in_event_name, in_event_time, in_event_result))
            
        connection.commit()
            
    except Exception as _ex:
        print("INFO  error", _ex)
    finally:
        if connection:
            connection.close()

pageNum = 2
while(True):
    url ="https://api.allsportdb.com/v3/calendar?page=" + str(pageNum) 
    response = requests.get(url, headers={'Authorization': 'Bearer 88d3ca9b-3004-484f-9669-a9112c115637'})
    jsonF = response.json()
    print(jsonF)
    if(jsonF == [] or pageNum == 15):
        break
    for i in jsonF:
        insertIntoDatabase(i["name"], i["dateFrom"], i["dateTo"], i["status"], i["sport"])
        print(i["name"], i["dateFrom"], i["dateTo"], i["status"], end="\n\n")
    pageNum += 1