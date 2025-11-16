# CLAUDE.md

このファイルは、Claude Codeなどの開発支援AIが参照するためのプロジェクト情報です。

## プロジェクト概要

Notion APIから取得したブロックデータ（JSON形式）をMarkdownに変換するCLIツール。
標準入力からJSONを読み込み、標準出力にMarkdownを出力する。

## アーキテクチャ

### ファイル構成

```
.
├── main.go           # エントリーポイント（標準入出力処理）
├── converter.go      # 変換ロジック（型定義 + Convert関数 + formatRichText関数）
├── converter_test.go # テストコード（16テストケース）
├── go.mod           # Goモジュール定義
└── README.md        # ユーザー向けドキュメント
```

### 主要な型定義

#### NotionResponse
Notion APIのレスポンス構造

```go
type NotionResponse struct {
    Object  string  `json:"object"`
    Results []Block `json:"results"`
}
```

#### Block
各ブロックの共通構造。ブロックタイプに応じたフィールドを持つ。

```go
type Block struct {
    Type             string
    Heading1         *Heading
    Heading2         *Heading
    Heading3         *Heading
    Paragraph        *Paragraph
    BulletedListItem *ListItem
    NumberedListItem *ListItem
    Code             *CodeBlock
}
```

#### RichText
テキストとアノテーション情報

```go
type RichText struct {
    PlainText   string
    Annotations Annotations
    Href        *string
}
```

### 主要な関数

#### `Convert(response NotionResponse) string`
NotionResponseを受け取り、Markdown文字列を返すメイン関数。
各ブロックタイプに応じて適切なMarkdownフォーマットに変換する。

#### `formatRichText(richTexts []RichText) string`
RichText配列を受け取り、アノテーション（太字、イタリック、コード、取り消し線）とリンクを処理してMarkdown文字列を返す。

## 開発方針

### t_wada式TDD

このプロジェクトは、t_wada式テスト駆動開発で実装されています。

**開発サイクル**:
1. **Red**: 失敗するテストを書く
2. **Green**: 最小限の実装でテストを通す
3. **Refactor**: コードを改善する

**実装の順序**:
1. 空のresults処理
2. heading_1, heading_2, heading_3
3. paragraph, bulleted_list_item, numbered_list_item
4. codeブロック
5. rich_textのアノテーション（bold, italic, code, strikethrough）
6. 複数アノテーションとリンク
7. 複数ブロックの処理
8. CLIツール実装
9. 実際のJSONでの動作確認

### テストケース（16個）

1. `TestConvertEmptyResults` - 空の配列
2. `TestConvertHeading1` - 見出し1
3. `TestConvertHeading2` - 見出し2
4. `TestConvertHeading3` - 見出し3
5. `TestConvertParagraph` - 段落
6. `TestConvertBulletedListItem` - 箇条書き
7. `TestConvertNumberedListItem` - 番号付きリスト
8. `TestConvertCodeBlock` - コードブロック
9. `TestConvertBoldAnnotation` - 太字
10. `TestConvertItalicAnnotation` - イタリック
11. `TestConvertCodeAnnotation` - インラインコード
12. `TestConvertStrikethroughAnnotation` - 取り消し線
13. `TestConvertMultipleAnnotations` - 複数アノテーション
14. `TestConvertLink` - リンク
15. `TestConvertMultipleRichTexts` - 複数のrich_text要素
16. `TestConvertMultipleBlocks` - 複数ブロック

## アノテーション処理の順序

`formatRichText`関数では、以下の順序でアノテーションを適用：

1. Code (`` `text` ``)
2. Bold (`**text**`)
3. Italic (`*text*`)
4. Strikethrough (`~~text~~`)
5. Link (`[text](url)`)

この順序により、`**bold**`と`*italic*`が同時に適用された場合、`***bold italic***`として正しくレンダリングされる。

## 今後の拡張可能性

現在サポートしていないが、将来追加可能な機能：

### ブロックタイプ
- `table` - テーブル
- `toggle` - 折りたたみ可能セクション
- `callout` - 注釈ボックス
- `quote` - 引用
- `divider` - 区切り線
- `image` - 画像
- `table_of_contents` - 目次

### 機能
- 子ブロックのサポート（ネストされたリストなど）
- 複数ファイルの一括変換
- エラーハンドリングの強化
- 変換オプション（Markdownスタイルのカスタマイズなど）

## コーディング規約

- テストコードには汎用的なテストデータを使用（実際のデータを避ける）
- 小さなステップでの実装（一度に一つの機能）
- テストファーストで開発
- コミット前に全テストが通ることを確認
