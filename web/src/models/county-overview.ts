import {DailyInstance} from './daily-instance';
import {MonthlyCountyDeath} from './monthly-deaths';


export interface CountyOverview{
  deaths: { [key: string]: MonthlyCountyDeath[]; };
  daily_cases: DailyInstance[];
  daily_deaths: DailyInstance[];
  daily_hospitalizations: DailyInstance[];
}
