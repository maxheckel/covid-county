import {Component, Input, OnInit} from '@angular/core';
import {CountyOverview} from '../../models/county-overview';
import {ChartDataSets, ChartType} from 'chart.js';
import {Color, Label} from 'ng2-charts';

@Component({
  selector: 'app-deaths-chart',
  templateUrl: './deaths-chart.component.html',
  styleUrls: ['./deaths-chart.component.scss']
})
export class DeathsChartComponent implements OnInit {
  @Input() county: CountyOverview;


  barChartData: ChartDataSets[] = [];

  barChartLabels: Label[] = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'];

  barChartOptions = {
    responsive: true,
    scales: {
      yAxes: [{
        ticks: {
          beginAtZero: true
        }
      }]
    }

  };

  barChartColors: Color[] = [
    {
      backgroundColor: '#235789'
    },
    {
      backgroundColor: '#C1292E'
    }
  ];

  barChartLegend = true;
  barChartPlugins = [];
  barChartType: ChartType = 'bar';

  constructor() {

  }

  ngOnInit(): void {
    Object.keys(this.county.deaths).forEach(key => {
      this.barChartData.push({
        label: key === '2020' ? key : 'Last 5 years avg',
        data: this.county.deaths[key].map(death => death.count)
      });
    });
  }

}
