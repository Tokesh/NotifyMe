# response = requests.get("https://api.allsportdb.com/v3/sports", headers={'Authorization': 'Bearer 88d3ca9b-3004-484f-9669-a9112c115637'})
# jsonF = response.json()
# print(jsonF)
import psycopg2
import requests
from config import host, user, password, db_name
def insertIntoDatabase():
    try:
        connection = psycopg2.connect(
            host=host,
            user = user,
            password = password
        )
        
        with connection.cursor() as cursor:
            #sql1 = "INSERT INTO events (events_id, event_name, event_time, event_result) VALUES (%s, %s, %s, %s)"
            cursor.execute(
                "Select subs_id from subscriptions where category2 = 'Basketball'"
            )
            print(len(cursor.fetchone()))
            # insert into events and dont forget to bind event_id with subscription_id
            #sql1 = "INSERT INTO events (event_name, event_time, event_result) VALUES (%s, %s, %s)"
            #cursor.execute(sql1, (in_event_name, in_event_time, in_event_result))
            
        connection.commit()
            
    except Exception as _ex:
        print("INFO  error", _ex)
    finally:
        if connection:
            connection.close()
insertIntoDatabase()
# import datetime

# from pytz import UTC
# yourdate = datetime.datetime.fromisoformat(str(input()))
# print(yourdate)
# print(datetime.datetime.astimezone(yourdate, UTC))