import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {County} from "../models/county";
import {Injectable} from "@angular/core";

@Injectable({
  providedIn: 'root'
})
export class ApiService{
  constructor(private client: HttpClient) {}

  getOverview():Observable<County[]>{
    return this.client.get<County[]>("/api/overview")
  }
}
