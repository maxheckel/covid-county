export interface MonthlyCountyDeath{
  month: number;
  county: string;
  year: number;
  count: number;
}

export interface CountyOverview{
  deaths: { [key: string]: MonthlyCountyDeath[]; };
}
