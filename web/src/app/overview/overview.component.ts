import {Component, OnInit} from '@angular/core';
import {ApiService} from "../../services/api.service";
import {County} from "../../models/county";
import {first} from "rxjs/operators";
import { ChartDataSets, ChartOptions } from 'chart.js';
import { Color, Label } from 'ng2-charts';

const trendingRatioSort = (a, b) => {
  return b.trending_ratio - a.trending_ratio;
};


@Component({
  selector: 'app-overview',
  templateUrl: './overview.component.html',
  styleUrls: ['./overview.component.scss']
})
export class OverviewComponent implements OnInit {

  public overviewData: County[];


  constructor(private service: ApiService) {
  }

  ngOnInit(): void {


    this.service.getOverview().pipe(first()).subscribe(value => this.overviewData = value);
  }

  trendingUp(): County[] {
    return this.overviewData && this.overviewData.filter(datum => datum.trending_direction === "Upwards").sort(trendingRatioSort);
  }

  trendingDownwards(): County[] {
    return this.overviewData && this.overviewData.filter(datum => datum.trending_direction === "Downwards").sort(trendingRatioSort);
  }

  steady(): County[] {
    return this.overviewData && this.overviewData.filter(datum => datum.trending_direction === "Steady").sort(trendingRatioSort);
  }

}
