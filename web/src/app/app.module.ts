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
import { CountyOverviewComponent } from './county-overview/county-overview.component';
import { DeathsChartComponent } from './deaths-chart/deaths-chart.component';
import { DailyCountsComponent } from './daily-counts/daily-counts.component';
import { CountyStatsComponent } from './county-stats/county-stats.component';



@NgModule({
  declarations: [
    AppComponent,
    OverviewComponent,
    CountyComponent,
    StatsComponent,
    CountyOverviewComponent,
    DeathsChartComponent,
    DailyCountsComponent,
    CountyStatsComponent
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
