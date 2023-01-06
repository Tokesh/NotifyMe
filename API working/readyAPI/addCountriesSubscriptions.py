import psycopg2
import datetime
import requests
from config import host, user, password, db_name


def insertIntoDatabase(in_sub_name, in_subs_code, category1, category2):
    try:
        connection = psycopg2.connect(
            host=host,
            user = user,
            password = password
        )
        
        with connection.cursor() as cursor:
            sql1 = "INSERT INTO subscriptions(subs_name, subs_code, category1, category2) VALUES (%s, %s,%s,%s)"
            cursor.execute(sql1, (in_sub_name, in_subs_code, category1, category2))
            
        connection.commit()
            
    except Exception as _ex:
        print("INFO  error", _ex)
    finally:
        if connection:
            connection.close()






pageNum = 1
while(pageNum < 30):
    url ="https://api.allsportdb.com/v3/countries?page=" + str(pageNum) 
    response = requests.get(url, headers={'Authorization': 'Bearer 88d3ca9b-3004-484f-9669-a9112c115637'})
    jsonF = response.json()
    print(jsonF)
    if(jsonF == []):
        break
    for i in jsonF:
        insertIntoDatabase(i["name"], i["code"], "Country", i["continent"])
        print(i["name"], end="\n\n")
    pageNum += 1