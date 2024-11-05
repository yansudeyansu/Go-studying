package models

import (
	"encoding/json"
	"os"
	"time"
)

// StudySession は学習記録の1セッションを表す
type StudySession struct {
	Date     time.Time `json:"date"`     // 学習日時
	Topic    string    `json:"topic"`    // 学習トピック
	Notes    string    `json:"notes"`    // 学習内容
	Duration int       `json:"duration"` // 学習時間（分）
}

// StudyTracker は学習記録の管理を行う
type StudyTracker struct {
	Sessions []StudySession `json:"sessions"`
	DataFile string
}

// NewStudyTracker は新しい学習記録管理インスタンスを生成する
func NewStudyTracker(dataFile string) *StudyTracker {
	tracker := &StudyTracker{
		Sessions: make([]StudySession, 0),
		DataFile: dataFile,
	}
	tracker.loadFromFile()
	return tracker
}

// AddSession は新しい学習セッションを追加する
func (st *StudyTracker) AddSession(topic string, notes string, duration int) error {
	session := StudySession{
		Date:     time.Now(),
		Topic:    topic,
		Notes:    notes,
		Duration: duration,
	}

	st.Sessions = append(st.Sessions, session)
	return st.saveToFile()
}

// loadFromFile はJSONファイルからデータを読み込む
func (st *StudyTracker) loadFromFile() error {
	if _, err := os.Stat(st.DataFile); os.IsNotExist(err) {
		return nil
	}

	data, err := os.ReadFile(st.DataFile)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &st.Sessions)
}

// saveToFile はデータをJSONファイルに保存する
func (st *StudyTracker) saveToFile() error {
	data, err := json.MarshalIndent(st.Sessions, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(st.DataFile, data, 0644)
}
