export interface County{
  county: string;
  averages: { [key: string]: number; };
  trending_direction: string;
  trending_ratio: number;
}
