// OpenAIResponse はOpenAI APIからのレスポンスを表す
type OpenAIResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int64    `json:"created"`
	Choices []Choice `json:"choices"`
}

// Choice はOpenAIからの応答の選択肢を表す
type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

// Message はチャットのメッセージを表す
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// OpenAIRequest はOpenAI APIへのリクエスト構造を表す
type OpenAIRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}

// StudySession は学習記録の1セッションを表す
type StudySession struct {
	Date     time.Time `json:"date"`     // 学習日時
	Topic    string    `json:"topic"`    // 学習トピック
	Notes    string    `json:"notes"`    // 学習メモ
	Duration int       `json:"duration"` // 学習時間（分）
	Summary  string    `json:"summary"`  // AI生成の要約
	Tags     []string  `json:"tags"`     // AI生成のタグ
}

// StudyTracker は学習記録の管理を行う
type StudyTracker struct {
	Sessions  []StudySession `json:"sessions"`
	DataFile  string         // 保存ファイルパス
	OpenAIKey string         // OpenAIのAPIキー
}

// NewStudyTracker は新しい学習記録管理インスタンスを生成する
func NewStudyTracker(dataFile string, openAIKey string) *StudyTracker {
	return &StudyTracker{
		DataFile:  dataFile,
		OpenAIKey: openAIKey,
	}
}

// AddSession は新しい学習セッションを追加する
func (st *StudyTracker) AddSession(topic string, notes string, duration int) error {
	// 実装内容...
}

// analyzeWithAI は学習内容のAI分析を実行する
func (st *StudyTracker) analyzeWithAI(topic, notes string) (string, []string, error) {
	// 実装内容...
}

// GetTopicSummary は指定トピックの学習サマリーを生成する
func (st *StudyTracker) GetTopicSummary(topic string) string {
	// 実装内容...
}

// GetStudyStats は学習統計情報を返す
func (st *StudyTracker) GetStudyStats() map[string]interface{} {
	// 実装内容...
}

// SearchByTags は指定タグに一致する学習セッションを検索する
func (st *StudyTracker) SearchByTags(searchTags []string) []StudySession {
	// 実装内容...
}

// saveToFile は学習セッションをJSONファイルに保存する
func (st *StudyTracker) saveToFile() error {
	// 実装内容...
}

// loadFromFile はJSONファイルから学習セッションを読み込む
func (st *StudyTracker) loadFromFile() error {
	// 実装内容...
}