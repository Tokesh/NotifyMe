import psycopg2
import datetime
import requests
from config import host, user, password, db_name
def insertIntoDatabase(in_event_name, in_event_result):
    try:
        connection = psycopg2.connect(
            host=host,
            user = user,
            password = password
        )
        
        with connection.cursor() as cursor:
            in_event_timestart = datetime.datetime.fromisoformat("2022-11-06T00:00:00")
            in_event_timeend = datetime.datetime.fromisoformat("2022-12-06T00:00:00")
            
            sql1 = "INSERT INTO events (event_name,event_timestart, event_timeend,event_result) VALUES (%s, %s, %s, %s) returning events_id"
            cursor.execute(sql1, (in_event_name, in_event_timestart, in_event_timeend, in_event_result))
            event_id = cursor.fetchone()[0]
            
            subs_id = "221"
            
            sql2 = "INSERT INTO event_sub(event_id, subs_id) VALUES (%s, %s)"
            cursor.execute(sql2, (event_id, subs_id))
            
        connection.commit()
            
    except Exception as _ex:
        print("INFO  error", _ex)
    finally:
        if connection:
            connection.close()


url ="https://flixster.p.rapidapi.com/movies/get-popularity"
response = requests.get(url, headers={'X-RapidAPI-Key': '8c844a0977msh8e2d2f08560e48bp18579djsnf3c52634ea2c',
    'X-RapidAPI-Host': 'flixster.p.rapidapi.com'})
jsonF = response.json()
print(jsonF)
if(jsonF == []): exit(0)
for i in jsonF["data"]["popularity"]:
    if(i == None): continue
    insertIntoDatabase(i["name"], i["posterImage"]["url"])
    print(i["name"], i["posterImage"]["url"])