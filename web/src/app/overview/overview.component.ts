import {Component, OnInit} from '@angular/core';
import {ApiService} from "../../services/api.service";
import {County} from "../../models/county";
import {first} from "rxjs/operators";
import { ChartDataSets, ChartOptions } from 'chart.js';
import { Color, Label } from 'ng2-charts';

const trendingRatioSort = (a, b) => {
  return a.county > b.county ? 1 : -1
};


@Component({
  selector: 'app-overview',
  templateUrl: './overview.component.html',
  styleUrls: ['./overview.component.scss']
})
export class OverviewComponent implements OnInit {

  public overviewData: County[];

  public searchText = "";

  public queryType = "cases";

  constructor(private service: ApiService) {
  }

  ngOnInit(): void {
    this.service.getOverview().pipe(first()).subscribe(value => this.overviewData = value);
  }



  trendingUp(): County[] {
    return this.overviewData && this.overviewData.filter(datum => datum.trending_direction === "Upwards").filter((county: County) => {
    return county.county.toLowerCase().search(this.searchText.toLowerCase().trim()) != -1;
  }).sort(trendingRatioSort);
  }

  trendingDownwards(): County[] {
    return this.overviewData && this.overviewData.filter(datum => datum.trending_direction === "Downwards").filter((county: County) => {
    return county.county.toLowerCase().search(this.searchText.toLowerCase().trim()) != -1;
  }).sort(trendingRatioSort);
  }

  steady(): County[] {
    return this.overviewData && this.overviewData.filter(datum => datum.trending_direction === "Steady").filter((county: County) => {
    return county.county.toLowerCase().search(this.searchText.toLowerCase().trim()) != -1;
  }).sort(trendingRatioSort);
  }

  search($event){
    this.searchText = $event.target.value;
  }

  setSearchType($event){
    this.queryType = $event.target.value;
    this.service.getOverview(this.queryType).pipe(first()).subscribe(value => this.overviewData = value)
  }

}
