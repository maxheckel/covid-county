import {Component, Input, OnInit} from '@angular/core';
import {County} from "../../models/county";
import {ChartDataSets} from "chart.js";
import {Color, Label} from "ng2-charts";

@Component({
  selector: 'app-county',
  templateUrl: './county.component.html',
  styleUrls: ['./county.component.scss']
})
export class CountyComponent implements OnInit {
  @Input() trending: string;
  @Input() county: County
  @Input() dataType: string;

  lineChartData: ChartDataSets[]

  lineChartLabels: Label[];

  lineChartOptions = {
    responsive: true,
    scales: {
      yAxes: [{
        ticks: {
          beginAtZero: true
        }
      }]
    }

  };

  lineChartColors: Color[];

  lineChartLegend = true;
  lineChartPlugins = [];
  lineChartType = 'line';
  constructor() {

  }

  trendingToColor():string{
    switch (this.trending){
      case 'downwards':
        return 'rgba(182, 211, 105, 0.28)';
      case 'steady':
        return 'rgba(255,255,0,0.28)';
      case 'upwards':
        return 'rgba(209, 0, 0, 0.28)';
    }
  }


  ngOnInit(): void {
    this.lineChartColors = [
      {
        borderColor: 'black',
        backgroundColor: this.trendingToColor(),
      },
    ]
    this.lineChartData = [
      {
        data: Object.values(this.county.averages),
        label: "7 day "+this.dataType+" average"
      }
    ]
    this.lineChartLabels = Object.keys(this.county.averages)
  }

}
