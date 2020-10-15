import {Component, Input, OnInit} from '@angular/core';
import {County} from "../../models/county";

@Component({
  selector: 'app-stats',
  templateUrl: './stats.component.html',
  styleUrls: ['./stats.component.scss']
})
export class StatsComponent implements OnInit {
  @Input() up: County[];
  @Input() steady: County[];
  @Input() down: County[];
  constructor() { }

  ngOnInit(): void {
  }

}
