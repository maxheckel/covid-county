import {Component, Input, OnInit} from '@angular/core';
import {ChartDataSets, ChartType} from 'chart.js';
import {Color, Label} from 'ng2-charts';
import {DailyInstance} from '../../models/daily-instance';

@Component({
  selector: 'app-daily-counts',
  templateUrl: './daily-counts.component.html',
  styleUrls: ['./daily-counts.component.scss']
})
export class DailyCountsComponent implements OnInit {
  @Input() instances: DailyInstance[];

  lineChartData: ChartDataSets[];

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
  lineChartType = 'bar';

  constructor() { }

  ngOnInit(): void {
    this.lineChartData = [
      {
        data: this.instances.map(val => val.count),
        label: this.instances[0].type
      }
    ];
    this.lineChartLabels = this.instances.map(val => val.date.split('T00')[0]);
    if (this.instances[0].type === 'Case'){
      this.lineChartColors = [
        {
          backgroundColor: '#12355B'
        }
      ];
    }
    if (this.instances[0].type === 'Death'){
      this.lineChartColors = [
        {
          backgroundColor: '#D72638'
        }
      ];
    }
    if (this.instances[0].type === 'Hospitalization'){
      this.lineChartColors = [
        {
          backgroundColor: '#420039'
        }
      ];
    }
  }

}
