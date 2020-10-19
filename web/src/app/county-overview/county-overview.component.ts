import { Component, OnInit } from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import {CountyOverview} from '../../models/county-overview';
import {ApiService} from '../../services/api.service';

@Component({
  selector: 'app-county-overview',
  templateUrl: './county-overview.component.html',
  styleUrls: ['./county-overview.component.scss']
})
export class CountyOverviewComponent implements OnInit {
  public county: CountyOverview;
  public name: string;
  constructor(private route: ActivatedRoute, private service: ApiService) {
    this.name = route.snapshot.paramMap.get('name');
    this.service.getCounty(this.name).subscribe(val => {
      this.county = val;
    });
  }

  ngOnInit(): void {

  }

}
