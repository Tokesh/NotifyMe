import psycopg2
import requests
from config import host, user, password, db_name
def insertIntoDatabase(in_subs_id, in_subs_name, in_subs_code, in_category1, in_category2):
    try:
        connection = psycopg2.connect(
            host=host,
            user = user,
            password = password
        )
        
        with connection.cursor() as cursor:
            # insert into events and dont forget to bind event_id with subscription_id
            #sql1 = "INSERT INTO events (event_name, event_time, event_result) VALUES (%s, %s, %s)"
            #cursor.execute(sql1, (in_event_name, in_event_time, in_event_result))
            sql1 = "INSERT INTO subscriptions (subs_id, subs_name, subs_code, category1, category2) VALUES (%s, %s, %s, %s, %s)"
            cursor.execute(sql1, (in_subs_id, in_subs_name, in_subs_code, in_category1, in_category2))
            
        connection.commit()
            
    except Exception as _ex:
        print("INFO  error", _ex)
    finally:
        if connection:
            connection.close()

response = requests.get("https://api.allsportdb.com/v3/sports?page=8", headers={'Authorization': 'Bearer 88d3ca9b-3004-484f-9669-a9112c115637'})
#sport - 7 pages, 'https://api.allsportdb.com/v3/sports?page=1'

jsonF = response.json()
for i in jsonF:
    insertIntoDatabase(i["id"], i["name"], i["name"], "sport", i["name"])
    print(int(i["id"]), i["name"], i["name"], "sport", i["name"], end="\n\n")
