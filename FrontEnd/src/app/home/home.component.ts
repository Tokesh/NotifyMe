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
  pastEvents:Event[] = []
  futureEvents:Event[] = []
  currentDate = new Date();
  constructor(private EventsService: EventsService) {
  }


  ngOnInit(): void {
    this.getUserId()
    this.getEvents()
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
      this.sortingEvents()
    })
  }
  sortingEvents(){
      for(let i = 0; i<this.eventss.length;i++){
          console.log(this.eventss[i])
          let eventTime = new Date(this.eventss[i].event_timestart)
          if(eventTime < this.currentDate){
            if(this.futureEvents.length >= 10) continue;
            this.futureEvents.push(this.eventss[i])
          }else{
            if(this.pastEvents.length >= 10) continue;
            this.pastEvents.push(this.eventss[i])
          }
      }
      console.log(this.futureEvents)
    console.log(this.pastEvents)
      console.log(this.currentDate)
  }

}
