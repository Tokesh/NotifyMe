export interface Product{
    id: number,
    name: string;
    description: string;
    price: number;
    filename: string;
    height: number;
    width: number;
    rating: number;
    count:number;
}

export interface Category{
    id: number;
    name: string;
}

export interface Shop{
    id: number;
    name: string;
    description: string;
}

export interface City{
    id: number;
    name: string;
}

export interface AuthToken {
    token: string;
}

export interface Event{
  events_id: string,
  event_name: string,
  event_timestart: string,
  event_timeend: string,
  event_result: string
}

export interface User{
  user_id: number,
  name: string,
  email: string,
  password: string,
  activationStatus: string
  status: number
}
