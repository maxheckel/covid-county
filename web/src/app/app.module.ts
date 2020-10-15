import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import {HttpClientModule} from "@angular/common/http";
import { OverviewComponent } from './overview/overview.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { CountyComponent } from './county/county.component';
import {ChartsModule} from "ng2-charts";
import { StatsComponent } from './stats/stats.component';

@NgModule({
  declarations: [
    AppComponent,
    OverviewComponent,
    CountyComponent,
    StatsComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    NgbModule,
    ChartsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
