import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {County} from "../models/county";
import {Injectable} from "@angular/core";
import {CountyOverview} from "../models/county-overview";

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  constructor(private client: HttpClient) {
  }

  getOverview(queryType?: string): Observable<County[]> {
    return this.client.get<County[]>("/api/overview?type=" + queryType)
  }

  getCounty(county: string): Observable<CountyOverview> {
    return this.client.get<CountyOverview>("/api/county/" + county);
  }
}
