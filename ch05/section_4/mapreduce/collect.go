package main

import "text/scanner"

// key:   token
// value: []scanner.Position
type intermediate map[string][]scanner.Position

// 모든 토큰의 Position 정보를 수집
func collect(in <-chan partial) intermediate {
	tokenPositions := make(intermediate, 10)
	for t := range in {
		tokenPositions.addPartial(t)
	}
	return tokenPositions
}

// partial 정보를 토큰별로 그루핑하여 intermediate 맵에 저장
func (m intermediate) addPartial(p partial) {
	positions, ok := m[p.token]
	if !ok {
		positions = make([]scanner.Position, 1)
	}

	m[p.token] = append(positions, p.Position)
}
