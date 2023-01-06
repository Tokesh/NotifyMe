import psycopg2
import datetime
import requests
from config import host, user, password, db_name

def insertIntoDatabase(in_event_name, in_event_timestart, in_event_timeend, in_event_result):
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
            
            subs_id = "222"
            
            sql2 = "INSERT INTO event_sub(event_id, subs_id) VALUES (%s, %s)"
            cursor.execute(sql2, (event_id, subs_id))
            
        connection.commit()
            
    except Exception as _ex:
        print("INFO  error", _ex)
    finally:
        if connection:
            connection.close()

client_id = "MzAxNjk0NTZ8MTY2Nzc1NzM2Mi4wNTQ3MjYx"
client_secret = "b593b19d341bada0f63a7b7fbdb09475d9a1f91cadf9d4db510acfb3571dbc34"
url = "https://api.seatgeek.com/2/events" + "?client_id=" + client_id + "&client_secret=" + client_secret
response = requests.get(url)
jsonF = response.json()
print(jsonF)
if(jsonF == []): exit(0)
for i in jsonF["events"]:
    if(i["type"] == "concert"):
        print(i["type"])
        for j in i["performers"]:
            #print(j["name"], i["datetime_utc"], i["url"], end="\n\n")
            insertIntoDatabase(j["name"],i["datetime_utc"],i["datetime_utc"],i["url"])
    #insertIntoDatabase(i["name"], i["posterImage"]["url"])
    #print(i["name"], i["posterImage"]["url"])