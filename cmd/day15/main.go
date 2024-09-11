package main

import (
	"bufio"
	"image"
	"io"
	"regexp"
	"strconv"
)

var regexSensor = regexp.MustCompile(`^Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)$`)

type sensorReport struct {
	ptSensor, ptBeacon image.Point
	dist               int
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -1 * i
}

func dist(one, other image.Point) int {
	return abs(one.X-other.X) + abs(one.Y-other.Y)
}

func parseReport(s string) (report sensorReport) {
	matches := regexSensor.FindStringSubmatch(s)
	report.ptSensor.X, _ = strconv.Atoi(matches[1])
	report.ptSensor.Y, _ = strconv.Atoi(matches[2])

	report.ptBeacon.X, _ = strconv.Atoi(matches[3])
	report.ptBeacon.Y, _ = strconv.Atoi(matches[4])

	report.dist = dist(report.ptSensor, report.ptBeacon)

	return
}

func scanYLine(reports []sensorReport, yAxis int) int {
	emptySpots := make(map[int]bool)
	write := func(n int) {
		_, ok := emptySpots[n]
		if ok {
			return
		}
		emptySpots[n] = true
	}
	for _, report := range reports {
		distToAxis := abs(report.ptSensor.Y - yAxis)

		if distToAxis > report.dist {
			continue
		}

		if report.ptBeacon.Y == yAxis {
			emptySpots[report.ptBeacon.X] = false
		}

		// we can take additional steps along the y-line
		// based on how close the y-line is to the sensor
		steps := report.dist - distToAxis
		if steps == 0 {
			write(report.ptSensor.X)
			continue
		}

		for i := 0; i < steps; i++ {
			write(report.ptSensor.X + i)
			write(report.ptSensor.X - i)
		}

	}

	acc := 0
	for _, isEmpty := range emptySpots {
		if isEmpty {
			acc++
		}
	}

	return acc
}

func parseReports(r io.Reader) (reports []sensorReport) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		reports = append(reports, parseReport(s.Text()))
	}
	return
}
