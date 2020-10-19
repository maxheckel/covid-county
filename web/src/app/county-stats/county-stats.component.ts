import {Component, Input, OnInit} from '@angular/core';
import {CountyOverview} from '../../models/county-overview';

@Component({
  selector: 'app-county-stats',
  templateUrl: './county-stats.component.html',
  styleUrls: ['./county-stats.component.scss']
})
export class CountyStatsComponent implements OnInit {
  @Input() county: CountyOverview;
  constructor() { }

  totalDeaths = 0;
  totalHospitalizations = 0;
  totalCases = 0;
  excessDeaths = 0;

  ngOnInit(): void {
    this.county.daily_cases.forEach(c => {
      this.totalCases += c.count;
    });
    this.county.daily_hospitalizations.forEach(c => {
      this.totalHospitalizations += c.count;
    });
    this.county.daily_deaths.forEach(c => {
      this.totalDeaths += c.count;
    });
    const now = new Date();
    console.log(now.getMonth());

    let total2020 = 0;
    let totalAvg = 0;
    for (let i = 1; i <= now.getMonth(); i++ ){
      total2020 += this.county.deaths['2020'].find(d => d.month === i).count;
      totalAvg += this.county.deaths['0'].find(d => d.month === i).count;
    }
    this.excessDeaths = total2020 - totalAvg;

  }

}
