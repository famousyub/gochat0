import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AppComponent } from './app.component';
import { WeatherComponent } from './weather/weather.component';
import {RouterModule} from '@angular/router';
import { ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { ApixuService } from "./apixu.service";
import {allAppRoutes} from './routes';
import { UvComponent } from './uv/uv.component';
import { CurrentWeatherComponent } from './current-weather/current-weather.component';
import { ForecastComponent } from './forecast/forecast.component';
import { HomePageComponent } from './home-page/home-page.component';
import { TopBarComponent } from './top-bar/top-bar.component';

import { MatSliderModule } from '@angular/material/slider';

import { MatButtonModule } from '@angular/material/button';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatInputModule } from '@angular/material/input';
import { MatTabsModule } from '@angular/material/tabs';
import { MatCardModule } from '@angular/material/card';
import { MatDividerModule } from '@angular/material/divider';
import { MatListModule } from '@angular/material/list';
//import { MatSliderModule } from '@angular/material/slider';


/*
import {
  ,
  ,
  ,
  ,
  ,
  ,

} from '@angular/material';*/
//import { HomePageComponent } from './home-page/home-page.component';
import { StoreModule } from '@ngrx/store';
import { locationReducer } from './location-reducer';
//import { TopBarComponent } from './top-bar/top-bar.component';
import { FormsModule } from '@angular/forms';
import { WeatherService } from './weather.service';
//import { CurrentWeatherComponent } from './current-weather/current-weather.component';
//import { ForecastComponent } from './forecast/forecast.component';
//import { UvComponent } from './uv/uv.component';
//import { AlertsComponent } from './alerts/alerts.component';
//import { HttpClientModule } from '@angular/common/http';


@NgModule({
  declarations: [
    AppComponent,
    WeatherComponent,
    UvComponent,
    CurrentWeatherComponent,
    ForecastComponent,
    HomePageComponent,
    TopBarComponent
  ],
  imports: [
    BrowserModule,
      RouterModule.forRoot(allAppRoutes),
      ReactiveFormsModule,
      HttpClientModule,
    BrowserAnimationsModule,
    MatButtonModule,
    MatToolbarModule,
    StoreModule.forRoot({
      loc: locationReducer
    }),
    FormsModule,
    MatInputModule,
    MatTabsModule,
    MatCardModule,
    HttpClientModule,
    MatDividerModule,
    MatListModule

  ],
  providers: [ApixuService,WeatherService],
  bootstrap: [AppComponent]
})
export class AppModule { }
