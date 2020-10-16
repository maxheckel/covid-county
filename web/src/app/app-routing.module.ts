import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {OverviewComponent} from "./overview/overview.component";
import {CountyOverviewComponent} from "./county-overview/county-overview.component";

const routes: Routes = [
  {
    path: "",
    component: OverviewComponent
  },
  {
    path: "county/:name",
    component: CountyOverviewComponent

  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
