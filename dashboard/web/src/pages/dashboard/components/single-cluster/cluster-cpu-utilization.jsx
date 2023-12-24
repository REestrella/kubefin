// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0
import React from "react";
import {
  Button,
  Box,
  AreaChart,
  Container,
  Header,
} from "@cloudscape-design/components";
import { Cluster } from "../model/cluster";
import { Pod } from "../model/pod";
import { Node } from "../model/node";
import { Memory } from "../model/memory";
import { CPU } from "../model/cpu";
import { keepTwoDecimal } from "../components-common";
import { ParseArrayIntoTimeSeriesData } from "../../../commons/common-components";


export default function ClusterCPUUtilization(props) {
  let cluster = props.cluster;
  if (cluster === undefined || cluster === null) {
    cluster = new Cluster(
      "-",
      "-",
      "-",
      "-",
      new Pod(0, 0, 0, 0),
      new Node(0, 0, 0, 0),
      new Memory(0, 0, 0, 0, [0, 0], [0, 0], [0, 0], [0, 0]),
      new CPU(0, 0, 0, 0, [0, 0], [0, 0], [0, 0], [0, 0])
    );
  }

  const requestedCPUArray = cluster.CPU.requestedCPUArray;
  const availableCPUArray = cluster.CPU.availableCPUArray;
  const systemReservedArray = cluster.CPU.systemReservedArray;
  const totalCPUArray = cluster.CPU.totalCPUArray;

  var cpuRequestedData = ParseArrayIntoTimeSeriesData(requestedCPUArray);
  var cpuAvailableData = ParseArrayIntoTimeSeriesData(availableCPUArray);
  var cpuSystemReservedData = ParseArrayIntoTimeSeriesData(systemReservedArray);

  var yMax = 1;
  totalCPUArray.map(function (item) {
    if (Number(item[1]) > yMax) {
      yMax = Number(item[1]);
    }
    return yMax;
  });

  var series = []
  if (cpuSystemReservedData.length !== 0) {
    series.push({
      title: "System reserved CPU",
      type: "area",
      data: cpuSystemReservedData,
      color: "#FFFF00",
      valueFormatter: function o(e) {
        return keepTwoDecimal(e) + " C";
      },
    })
  }
  if (cpuRequestedData.length !== 0) {
    series.push({
      title: "Requested CPU",
      type: "area",
      data: cpuRequestedData,
      color: "#FF0000",
      valueFormatter: function o(e) {
        return keepTwoDecimal(e) + " C";
      },
    })
  }
  if (cpuAvailableData.length !== 0) {
    series.push({
      title: "Available CPU",
      type: "area",
      color: "#0000CD",
      data: cpuAvailableData,
      valueFormatter: function o(e) {
        return keepTwoDecimal(e) + " C";
      },
    })
  };

  return (
    <Container
      className="custom-dashboard-container"
      header={<Header variant="h2">CPU utilization</Header>}
    >
      <AreaChart
        series={series}
        xDomain={[cpuAvailableData[0].x, cpuAvailableData[cpuAvailableData.length - 1].x]}
        yDomain={[0, yMax * 1.2]}
        i18nStrings={{
          filterLabel: "Filter displayed data",
          filterPlaceholder: "Filter data",
          filterSelectedAriaLabel: "selected",
          detailPopoverDismissAriaLabel: "Dismiss",
          legendAriaLabel: "Legend",
          chartAriaRoleDescription: "line chart",
          detailTotalLabel: "Total",
          xTickFormatter: (e) =>
            e.toLocaleDateString("en-US", {
                month: "short",
                day: "numeric",
                hour: "numeric",
                minute: "numeric",
                second: "numeric",
                hour12: !1,
              })
              .split(",")
              .join("\n"),
          yTickFormatter: function o(e) {
            return keepTwoDecimal(e) + " C"
          },
        }}
        ariaLabel="Stacked area chart"
        errorText="Error loading data."
        height={300}
        hideFilter
        loadingText="Loading chart"
        recoveryText="Retry"
        xScaleType="time"
        xTitle={"Time(" + Intl.DateTimeFormat().resolvedOptions().timeZone + ")"}
        yTitle="CPU Cores (vCPU)"
        empty={
          <Box textAlign="center" color="inherit">
            <b>No data available</b>
            <Box variant="p" color="inherit">
              There is no data available
            </Box>
          </Box>
        }
        noMatch={
          <Box textAlign="center" color="inherit">
            <b>No matching data</b>
            <Box variant="p" color="inherit">
              There is no matching data to display
            </Box>
            <Button>Clear filter</Button>
          </Box>
        }
      />
    </Container>
  );
}
