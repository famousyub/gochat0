import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { WeatherComponent } from './weather/weather.component';
import {RouterModule} from '@angular/router';
import { ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { ApixuService } from "./apixu.service";
import {allAppRoutes} from './routes';

@NgModule({
  declarations: [
    AppComponent,
    WeatherComponent
  ],
  imports: [
    BrowserModule,
      RouterModule.forRoot(allAppRoutes),
      ReactiveFormsModule,
      HttpClientModule

  ],
  providers: [ApixuService],
  bootstrap: [AppComponent]
})
export class AppModule { }
