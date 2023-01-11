import smtplib
import os
import psycopg2
import datetime
import requests
from config import host, user, password, db_name
from email.message import EmailMessage

def insertIntoDatabase():
    try:
        connection = psycopg2.connect(
            host=host,
            user = user,
            password = password
        )
        
        with connection.cursor() as cursor:
            sql1 = """
                SELECT user_subscription.user_id,event_name, event_timestart, u.user_email FROM user_subscription
                    LEFT JOIN subscriptions ON user_subscription.subs_id = subscriptions.subs_id
                    LEFT JOIN event_sub ON subscriptions.subs_id = event_sub.subs_id
                    LEFT JOIN events e ON event_sub.event_id = e.events_id
                    LEFT JOIN users u ON u.user_id = user_subscription.user_id
                    where event_timestart >= CURRENT_DATE and event_timestart < CURRENT_DATE -INTERVAL '-1 DAY'
                    group by user_subscription.user_id, event_name, event_timestart, u.user_email;
            """
            
            cursor.execute(sql1)
            row = cursor.fetchone()
            print(row)
            prevId = row[0]
            email = row[3]
            
            events = []
            while(row != None):
                if(row[0] != prevId):
                    print(prevId, events, email)
                    print(send_mail(email, events))
                    prevId = row[0]
                    events = []
                    email = row[3]
                eventName = row[1]
                events.append(eventName)
                row = cursor.fetchone()
            if(row == None): 
                print(prevId, events, email)
                print(send_mail(email, events))
        
        connection.commit()
            
    except Exception as _ex:
        print("INFO  error", _ex)
    finally:
        if connection:
            connection.close()


def send_mail(receiver, events):
    
    sender = "tokesh.api@mail.ru"
    password = "sJC9rDkmY8y6xhjAeUt5"
    server = smtplib.SMTP("smtp.mail.ru", 587)
    print("Подключился к SMTP", "\n")
    server.starttls()
    msg = "Good morning! Today events: \n"
    for i in range(len(events)):
        msg += f"{i+1}. {events[i]} \n"
    print("Вот сообщение которое отправлю ", msg)

    try:
        message = 'Subject: {}\n\n{}'.format("NotifyMe today events", msg)
        print("Начинаю попытку", "\n")
        server.login(sender,password)
        print("Залогинился", "\n")
        server.sendmail(sender,receiver, message)
        
        return "The message was sent successfuly"
    except Exception as _ex:
        return f"{_ex}, Check login and password"


def main():
    insertIntoDatabase()
    #print(send_mail())


if __name__ == '__main__':
    main()
