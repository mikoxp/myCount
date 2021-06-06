import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import {ChartDataSets} from "chart.js";
import {Color, Label} from "ng2-charts";

@Component({
  selector: 'app-steps',
  templateUrl: './steps.component.html',
  styleUrls: ['./steps.component.css']
})
export class StepsComponent implements OnInit {

  constructor(private http: HttpClient) { 
    this.http.get<any>('/api/steps').subscribe(res=>{
      console.log(res);
    })
  }

  ngOnInit(): void {
  }
  lineChartData: ChartDataSets[] = [
    { data: [85, 72, 78, 75, 77, 75], label: 'Crude oil prices' },
  ];

  lineChartLabels: Label[] = ['January', 'February', 'March', 'April', 'May', 'June'];

  lineChartOptions = {
    responsive: true,
  };

  lineChartColors: Color[] = [
    {
      borderColor: 'black',
      backgroundColor: 'rgba(255,255,0,0.28)',
    },
  ];

  lineChartLegend = true;
  lineChartPlugins = [];
  lineChartType = 'line';

}
