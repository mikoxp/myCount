import { HttpClient } from '@angular/common/http';
import { analyzeAndValidateNgModules } from '@angular/compiler';
import { Component, OnInit } from '@angular/core';
import {ChartDataSets} from "chart.js";
import {Color, Label} from "ng2-charts";

@Component({
  selector: 'app-steps',
  templateUrl: './steps.component.html',
  styleUrls: ['./steps.component.css']
})
export class StepsComponent implements OnInit {

  private label:Label[]=[];
  private data:number[]=[];

  lineChartData: ChartDataSets[] = [];

  lineChartLabels: Label[] = [];

  lineChartOptions = {
    responsive: true,
  };

  lineChartColors: Color[] = [
    {
      borderColor: 'black',
      backgroundColor: 'rgba(0, 102, 204,0.28)',
    },
  ];

  lineChartLegend = false;
  lineChartPlugins = [];
  lineChartType = 'line';

  constructor(private http: HttpClient) { 
    this.http.get<any>('/api/steps').subscribe((res:any[])=>{
      res.forEach(element => {
        this.label.push(element.day);
        this.data.push(element.stepsNumber);
      });
      console.log(this.label,this.data);
      this.lineChartData=[{ data: this.data, label: 'Steps' }]
      this.lineChartLabels=this.label;
    })
  }

  ngOnInit(): void {
  }

}
