import {Component, NgZone, OnInit} from '@angular/core';
import {EventsService} from "../events.service";
import {Event, User} from "../models";
@Component({
  selector: 'home-page',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']

})
export class HomeComponent implements OnInit {
  eventss:Event[] = [];
  user_id = 0;
  token = localStorage.getItem('token');
  constructor(private EventsService: EventsService) {
  }


  ngOnInit(): void {
    this.getUserId()
    this.getEvents()
    console.log(this.eventss)
  }
  getUserId(){
    let token = localStorage.getItem('token');
    this.EventsService.getUserIdHelp(token || "").subscribe((data:User) => {
      this.user_id = data.user_id
      this.getEvents()
    })

  }
  getEvents() {
    this.EventsService.getEventHelp(this.user_id).subscribe((data) => {

      // @ts-ignore
      this.eventss = data['events'];

      console.log(this.eventss)
    })
  }

}
