import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
@Injectable({
  providedIn: 'root'
})
export class ApixuService {

  constructor(private http: HttpClient) { }
  /*
  http://api.weatherstack.com/current
      ? access_key = YOUR_ACCESS_KEY
      & query = New York
  */
  api_key = "f383b0dde5de6382d0375f835e4d8056";
  getWeather(location){
      return this.http.get(
        'http://api.weatherstack.com/current?access_key='+this.api_key + '&query=' + location

        //  'https://api.apixu.com/v1/current.json?key='+ this.api_key + '&q=' + location
      );
  }
}
