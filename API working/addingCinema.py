import psycopg2
import datetime
import requests
from config import host, user, password, db_name
def insertIntoDatabase(in_events_id, in_event_name, in_event_time, in_event_result, in_subs_id, in_unique_id):
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
            #sql3
            
        connection.commit()
            
    except Exception as _ex:
        print("INFO  error", _ex)
    finally:
        if connection:
            connection.close()

url ="https://flixster.p.rapidapi.com/movies/get-upcoming"
response = requests.get(url, headers={'X-RapidAPI-Key': '8c844a0977msh8e2d2f08560e48bp18579djsnf3c52634ea2c', 'X-RapidAPI-Host':'flixster.p.rapidapi.com'})
jsonF = response.json()
print(jsonF)
for i in jsonF["data"]["upcoming"]:
    insertIntoDatabase(i["id"], i["name"], i["dateFrom"], "", i["sportId"])
    print(i["id"], i["name"], i["dateFrom"], "", i["sportId"], end="\n\n")