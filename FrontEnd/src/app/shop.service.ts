import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { AuthToken, City, Shop } from './models';
import { Observable } from 'rxjs';
import { HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class ShopService {

  BASE_URL = 'http://localhost:8080';

  constructor(private http: HttpClient) {

  }

  login(username: string, password: string): Observable<AuthToken> {
    return this.http.post<AuthToken>(`${this.BASE_URL}/user/login`, {
      'username':username,
      'user_password':password
    });
  }

  getCities(): Observable<City[]> {
    return this.http.get<City[]>(`${this.BASE_URL}/api/cities/`);
  }

  getShops(): Observable<Shop[]> {
    return this.http.get<Shop[]>(`${this.BASE_URL}/api/shops/`);
  }

  getShopsByCity(c_id: number): Observable<Shop[]> {
    return this.http.get<Shop[]>(`${this.BASE_URL}/api/cities/${c_id}/`);
  }

}
