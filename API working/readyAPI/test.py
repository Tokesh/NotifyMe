import requests
client_id = "MzAxNjk0NTZ8MTY2Nzc1NzM2Mi4wNTQ3MjYx"
client_secret = "b593b19d341bada0f63a7b7fbdb09475d9a1f91cadf9d4db510acfb3571dbc34"
url = "https://api.seatgeek.com/2/performers" + "?client_id=" + client_id + "&client_secret=" + client_secret
response = requests.get(url)
jsonF = response.json()
print(jsonF)
if(jsonF == []): exit(0)