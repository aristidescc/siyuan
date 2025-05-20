// SiYuan - Refactor your thinking
// Copyright (c) 2020-present, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package conf

import "github.com/siyuan-note/siyuan/kernel/util"

type Editor struct {
	AllowHTMLBLockScript            bool           `json:"allowHTMLBLockScript"`            // Allow execution of scripts within HTML blocks
	FontSize                        int            `json:"fontSize"`                        // Font size
	FontSizeScrollZoom              bool           `json:"fontSizeScrollZoom"`              // Whether font size supports mouse wheel zoom
	FontFamily                      string         `json:"fontFamily"`                      // Font family
	CodeSyntaxHighlightLineNum      bool           `json:"codeSyntaxHighlightLineNum"`      // Whether to display line numbers in code blocks
	CodeTabSpaces                   int            `json:"codeTabSpaces"`                   // Number of spaces to convert tabs to in code blocks. Set to 0 to disable conversion
	CodeLineWrap                    bool           `json:"codeLineWrap"`                    // Whether to enable line wrapping in code blocks
	CodeLigatures                   bool           `json:"codeLigatures"`                   // Whether to enable ligatures in code blocks
	DisplayBookmarkIcon             bool           `json:"displayBookmarkIcon"`             // Whether to display bookmark icons
	DisplayNetImgMark               bool           `json:"displayNetImgMark"`               // Whether to display network image badges
	GenerateHistoryInterval         int            `json:"generateHistoryInterval"`         // History generation interval in minutes
	HistoryRetentionDays            int            `json:"historyRetentionDays"`            // Number of days to retain history
	Emoji                           []string       `json:"emoji"`                           // Frequently used emojis
	VirtualBlockRef                 bool           `json:"virtualBlockRef"`                 // Whether to enable virtual references
	VirtualBlockRefExclude          string         `json:"virtualBlockRefExclude"`          // Virtual reference keyword exclusion list
	VirtualBlockRefInclude          string         `json:"virtualBlockRefInclude"`          // Virtual reference keyword inclusion list
	BlockRefDynamicAnchorTextMaxLen int            `json:"blockRefDynamicAnchorTextMaxLen"` // Maximum length of dynamic anchor text for block references
	PlantUMLServePath               string         `json:"plantUMLServePath"`               // PlantUML server address
	FullWidth                       bool           `json:"fullWidth"`                       // Whether to use maximum width
	KaTexMacros                     string         `json:"katexMacros"`                     // KaTex macro definitions
	ReadOnly                        bool           `json:"readOnly"`                        // Read-only mode
	EmbedBlockBreadcrumb            bool           `json:"embedBlockBreadcrumb"`            // Whether to display breadcrumbs for embedded blocks
	ListLogicalOutdent              bool           `json:"listLogicalOutdent"`              // List logical reverse indentation
	ListItemDotNumberClickFocus     bool           `json:"listItemDotNumberClickFocus"`     // Focus on list item marker when clicked
	FloatWindowMode                 int            `json:"floatWindowMode"`                 // Floating window trigger mode, 0: cursor hover, 1: Ctrl + hover, 2: no floating window
	DynamicLoadBlocks               int            `json:"dynamicLoadBlocks"`               // Dynamic block count, configurable range [48, 1024]
	Justify                         bool           `json:"justify"`                         // Whether to enable text justification
	RTL                             bool           `json:"rtl"`                             // Whether to display text right-to-left
	Spellcheck                      bool           `json:"spellcheck"`                      // Whether to enable spell checking
	OnlySearchForDoc                bool           `json:"onlySearchForDoc"`                // Whether to enable [[ to search document blocks only
	BacklinkExpandCount             int            `json:"backlinkExpandCount"`             // Default number of expanded backlinks
	BackmentionExpandCount          int            `json:"backmentionExpandCount"`          // Default number of expanded back mentions
	BacklinkContainChildren         bool           `json:"backlinkContainChildren"`         // Whether backlinks include child blocks in calculations
	Markdown                        *util.Markdown `json:"markdown"`                        // Markdown configuration
}

const (
	MinDynamicLoadBlocks = 48
)

func NewEditor() *Editor {
	return &Editor{
		FontSize:                        16,
		FontSizeScrollZoom:              false,
		CodeSyntaxHighlightLineNum:      false,
		CodeTabSpaces:                   0,
		CodeLineWrap:                    false,
		CodeLigatures:                   false,
		DisplayBookmarkIcon:             true,
		DisplayNetImgMark:               true,
		GenerateHistoryInterval:         10,
		HistoryRetentionDays:            30,
		Emoji:                           []string{},
		VirtualBlockRef:                 false,
		BlockRefDynamicAnchorTextMaxLen: 96,
		PlantUMLServePath:               "https://www.plantuml.com/plantuml/svg/~1",
		FullWidth:                       true,
		KaTexMacros:                     "{}",
		ReadOnly:                        false,
		EmbedBlockBreadcrumb:            false,
		ListLogicalOutdent:              false,
		ListItemDotNumberClickFocus:     true,
		FloatWindowMode:                 0,
		DynamicLoadBlocks:               192,
		Justify:                         false,
		RTL:                             false,
		BacklinkExpandCount:             8,
		BackmentionExpandCount:          -1,
		BacklinkContainChildren:         true,
		Markdown:                        util.MarkdownSettings,
	}
}
