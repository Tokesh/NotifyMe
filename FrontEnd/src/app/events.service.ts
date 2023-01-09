import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {AuthToken, City, Shop, User, Event} from "./models";

@Injectable({
  providedIn: 'root'
})
export class EventsService {
  BASE_URL = 'http://localhost:8080';

  constructor(private http: HttpClient) {

  }

  getEventHelp(event_id: Number): Observable<Event[]> {
    return this.http.get<Event[]>(`${this.BASE_URL}/event/${event_id}`);
  }
  getUserIdHelp(token: String): Observable<User> {
    return this.http.post <User>(`${this.BASE_URL}/user/`, {
      'token':token
    });
  }


}
