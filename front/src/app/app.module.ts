import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import {ChartsModule} from "ng2-charts";
import { AppRoutingModule } from './app-routing.module';
import { StepsComponent } from './steps/steps.component';


@NgModule({
  declarations: [
    AppComponent,
    StepsComponent
  ],
  imports: [
    BrowserModule,
    ChartsModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }