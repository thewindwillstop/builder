package controller

import (
	"fmt"
	"strings"
	"github.com/goplus/builder/spx-backend/internal/svggen"
)

// ThemeType represents different SVG generation themes
type ThemeType string

const (
	ThemeNone      ThemeType = ""
	ThemeCartoon   ThemeType = "cartoon"
	ThemeRealistic ThemeType = "realistic"
	ThemeMinimal   ThemeType = "minimal"
	ThemeFantasy   ThemeType = "fantasy"
	ThemeRetro     ThemeType = "retro"
	ThemeScifi     ThemeType = "scifi"
	ThemeNature    ThemeType = "nature"
	ThemeBusiness  ThemeType = "business"
)

// ThemeInfo represents detailed information about a theme
type ThemeInfo struct {
	ID                  ThemeType       `json:"id"`
	Name                string          `json:"name"`
	Description         string          `json:"description"`
	Prompt              string          `json:"prompt"`
	RecommendedProvider svggen.Provider `json:"recommended_provider"`
	PreviewURL          string          `json:"preview_url"`
}

// ThemePrompts maps each theme to its corresponding prompt enhancement
var ThemePrompts = map[ThemeType]string{
	ThemeCartoon:   "必须使用卡通风格，必须色彩鲜艳丰富，必须可爱有趣，严格使用简单几何形状，强制使用明亮饱和的色彩，禁止写实细节",
	ThemeRealistic: "必须使用写实风格，严格要求高度细节化，强制逼真效果，必须专业高质量渲染，禁止卡通化或简化元素",
	ThemeMinimal:   "必须使用极简风格，严格限制元素数量，强制使用干净线条和几何形状，严格使用黑白或单色调，禁止复杂装饰",
	ThemeFantasy:   "必须使用奇幻魔法风格，强制添加神秘魔法元素，严格使用梦幻色彩，必须包含超自然效果，禁止现实主义元素",
	ThemeRetro:     "必须使用复古怀旧风格，严格遵循经典老式美学，强制使用怀旧色调和设计元素，禁止现代化元素",
	ThemeScifi:     "必须使用科幻未来风格，强制添加科技元素，严格使用霓虹和金属色彩，必须包含未来感设计，禁止传统元素",
	ThemeNature:    "必须使用自然有机风格，严格使用自然元素和植物，强制使用大地色调和绿色系，禁止人工几何元素",
	ThemeBusiness:  "必须使用商务专业风格，严格保持企业形象，强制使用现代简洁设计，必须专业精致，禁止卡通或娱乐元素",
}

// ThemeNames maps each theme to its Chinese name
var ThemeNames = map[ThemeType]string{
	ThemeNone:      "无主题",
	ThemeCartoon:   "卡通风格",
	ThemeRealistic: "写实风格",
	ThemeMinimal:   "极简风格",
	ThemeFantasy:   "奇幻风格",
	ThemeRetro:     "复古风格",
	ThemeScifi:     "科幻风格",
	ThemeNature:    "自然风格",
	ThemeBusiness:  "商务风格",
}

// ThemePreviewPrompts maps each theme to its preview generation prompt
var ThemePreviewPrompts = map[ThemeType]string{
	ThemeNone:      "Create a simple, clean design preview showing basic shapes and neutral colors",
	ThemeCartoon:   "Create a colorful cartoon-style preview with bright colors, simple geometric shapes, cute characters or objects, playful and fun atmosphere",
	ThemeRealistic: "Create a realistic, detailed preview with photographic quality, fine details, natural lighting, professional high-quality rendering",
	ThemeMinimal:   "Create a minimalist preview with clean lines, geometric shapes, lots of white space, limited color palette (black, white, or single accent color)",
	ThemeFantasy:   "Create a fantasy-themed preview with magical elements, mystical creatures, enchanted atmosphere, dreamy colors and supernatural effects",
	ThemeRetro:     "Create a retro/vintage preview with classic design elements, nostalgic color schemes, old-fashioned aesthetics, vintage typography or patterns",
	ThemeScifi:     "Create a sci-fi preview with futuristic elements, neon colors, metallic surfaces, technological devices, space or cyberpunk atmosphere",
	ThemeNature:    "Create a nature-themed preview with organic elements, plants, natural textures, earth tones and green colors, outdoor scenery",
	ThemeBusiness:  "Create a professional business preview with modern corporate design, clean professional aesthetics, business icons or charts",
}

// ThemeDescriptions maps each theme to its description
var ThemeDescriptions = map[ThemeType]string{
	ThemeNone:      "不应用任何特定主题风格",
	ThemeCartoon:   "色彩鲜艳的卡通风格，适合可爱有趣的内容",
	ThemeRealistic: "高度写实的风格，细节丰富逼真",
	ThemeMinimal:   "极简主义风格，简洁干净的设计",
	ThemeFantasy:   "充满魔法和超自然元素的奇幻风格",
	ThemeRetro:     "怀旧复古风格，经典老式美学",
	ThemeScifi:     "未来科技风格，充满科幻元素",
	ThemeNature:    "自然有机风格，使用自然元素和大地色调",
	ThemeBusiness:  "专业商务风格，现代企业形象",
}

// ThemeProviders maps each theme to its recommended provider
var ThemeProviders = map[ThemeType]svggen.Provider{
	ThemeNone:      svggen.ProviderOpenAI,   // Default provider
	ThemeCartoon:   svggen.ProviderRecraft, // Recraft excels at cartoon styles
	ThemeRealistic: svggen.ProviderRecraft, // Recraft is good for realistic styles
	ThemeMinimal:   svggen.ProviderSVGIO,   // SVGIO works well for minimal styles
	ThemeFantasy:   svggen.ProviderRecraft, // Recraft handles fantasy well
	ThemeRetro:     svggen.ProviderRecraft, // Recraft for retro styles
	ThemeScifi:     svggen.ProviderRecraft, // Recraft for sci-fi
	ThemeNature:    svggen.ProviderRecraft, // Recraft for nature themes
	ThemeBusiness:  svggen.ProviderRecraft, // Recraft for business styles
}

// IsValidTheme checks if the given theme is valid
func IsValidTheme(theme ThemeType) bool {
	if theme == ThemeNone {
		return true
	}
	_, exists := ThemePrompts[theme]
	return exists
}

// GetThemePromptEnhancement returns the prompt enhancement for a given theme
func GetThemePromptEnhancement(theme ThemeType) string {
	if theme == ThemeNone {
		return ""
	}
	return ThemePrompts[theme]
}

// GetThemeRecommendedProvider returns the recommended provider for a given theme
func GetThemeRecommendedProvider(theme ThemeType) svggen.Provider {
	if provider, exists := ThemeProviders[theme]; exists {
		return provider
	}
	return svggen.ProviderSVGIO // Default fallback
}

// ApplyThemeToPrompt applies theme enhancement to the original prompt
func ApplyThemeToPrompt(originalPrompt string, theme ThemeType) string {
	if theme == ThemeNone {
		return originalPrompt
	}

	themeEnhancement := GetThemePromptEnhancement(theme)
	if themeEnhancement == "" {
		return originalPrompt
	}

	return fmt.Sprintf("%s，%s", originalPrompt, themeEnhancement)
}

// GetAvailableThemes returns all available themes
func GetAvailableThemes() []ThemeType {
	return []ThemeType{
		ThemeNone,
		ThemeCartoon,
		ThemeRealistic,
		ThemeMinimal,
		ThemeFantasy,
		ThemeRetro,
		ThemeScifi,
		ThemeNature,
		ThemeBusiness,
	}
}

// ThemePreviewURLs maps each theme to its preview image URL
var ThemePreviewURLs = map[ThemeType]string{
	ThemeNone:      "kodo://goplus-builder-usercontent-test/files/d39c908e01500cc777fbc98be99eae4e-54273.jpg",
	ThemeCartoon:   "kodo://goplus-builder-usercontent-test/files/0a5812620550338caed6c6f3d7cc858d-89145.svg",
	ThemeRealistic: "kodo://goplus-builder-usercontent-test/files/f11ef23e709f1e23f7aafb2a04923208-336687.svg",
	ThemeMinimal:   "kodo://goplus-builder-usercontent-test/files/9fa5c9b82e17d9833934dae6e233a13c-56750.svg",
	ThemeFantasy:   "kodo://goplus-builder-usercontent-test/files/0e21bf8fed236c7d9d3fe0f472429629-153495.svg",
	ThemeRetro:     "kodo://goplus-builder-usercontent-test/files/0185cc215cb268378e2a67427ad854a1-180335.jpg",
	ThemeScifi:     "kodo://goplus-builder-usercontent-test/files/58cdec62945556dc327dfc8117ea1ab9-142801.jpg",
	ThemeNature:    "kodo://goplus-builder-usercontent-test/files/cd7768e5fb3ef6f0bd2e2125f9298da7-181247.jpg",
	ThemeBusiness:  "kodo://goplus-builder-usercontent-test/files/9918818066ba6ad7478a5f79870adad2-122571.jpg",
}

// GetThemePreviewURL returns the preview image URL for a specific theme
func GetThemePreviewURL(theme ThemeType) string {
	if url, exists := ThemePreviewURLs[theme]; exists {
		return url
	}
	return "https://img.recraft.ai/FalrtQiAGrRlsDJ8wugqlVoQPHL1eSLsaZhy0AHTuB4/rs:fit:1024:1024:0/raw:1/plain/abs://external/images/bfd6dd23-a744-409e-931d-b8c38409fa41" // Default fallback
}

// GetThemeInfo returns detailed information for a specific theme
func GetThemeInfo(theme ThemeType) ThemeInfo {
	return ThemeInfo{
		ID:                  theme,
		Name:                ThemeNames[theme],
		Description:         ThemeDescriptions[theme],
		Prompt:              ThemePrompts[theme],
		RecommendedProvider: GetThemeRecommendedProvider(theme),
		PreviewURL:          GetThemePreviewURL(theme),
	}
}

// GetAllThemesInfo returns detailed information for all themes
func GetAllThemesInfo() []ThemeInfo {
	themes := GetAvailableThemes()
	result := make([]ThemeInfo, len(themes))

	for i, theme := range themes {
		result[i] = GetThemeInfo(theme)
	}

	return result
}




// PromptAnalysis represents the analysis result of a user prompt
type PromptAnalysis struct {
	Type       string   `json:"type"`       // "animal", "object", "scene", "character", etc.
	Emotion    string   `json:"emotion"`    // "cute", "serious", "scary", "neutral", etc.
	Complexity string   `json:"complexity"` // "simple", "medium", "complex"
	Keywords   []string `json:"keywords"`   // Extracted keywords
}

// QualityPrompts defines quality enhancement prompts for different complexity levels
var QualityPrompts = map[string]string{
	"simple":  "high quality vector art, clean lines, vibrant colors",
	"medium":  "high quality vector art, detailed illustration, professional design, rich colors",
	"complex": "high quality vector art, intricate details, professional illustration, sophisticated design, rich color palette",
}

// StylePrompts defines style enhancement prompts for different types
var StylePrompts = map[string]string{
	"animal":    "cute and friendly style, appealing character design",
	"object":    "modern and functional style, clear visual representation",
	"scene":     "atmospheric and immersive style, detailed environment",
	"character": "expressive and memorable style, distinctive personality",
	"nature":    "organic and natural style, harmonious composition",
	"building":  "architectural and structured style, clear geometric forms",
	"food":      "appetizing and colorful style, fresh appearance",
	"vehicle":   "sleek and dynamic style, modern design",
	"abstract":  "artistic and creative style, imaginative composition",
	"default":   "appealing and well-designed style, visual clarity",
}

// TechnicalPrompts defines technical requirements for SVG generation
var TechnicalPrompts = map[ThemeType]string{
	ThemeCartoon:   "SVG format, scalable graphics, cartoon-optimized, bright color palette",
	ThemeRealistic: "SVG format, scalable graphics, realistic rendering, detailed textures",
	ThemeMinimal:   "SVG format, scalable graphics, minimal design, geometric shapes",
	ThemeFantasy:   "SVG format, scalable graphics, magical effects, fantasy elements",
	ThemeRetro:     "SVG format, scalable graphics, vintage aesthetics, retro styling",
	ThemeScifi:     "SVG format, scalable graphics, futuristic design, tech elements",
	ThemeNature:    "SVG format, scalable graphics, natural textures, organic forms",
	ThemeBusiness:  "SVG format, scalable graphics, professional design, corporate styling",
	ThemeNone:      "SVG format, scalable graphics, web optimized",
}

// AnalyzePromptType analyzes the user prompt to determine its type, emotion, and complexity
func AnalyzePromptType(prompt string) PromptAnalysis {
	prompt = strings.ToLower(prompt)
	
	// Initialize analysis
	analysis := PromptAnalysis{
		Type:       "default",
		Emotion:    "neutral",
		Complexity: "simple",
		Keywords:   []string{},
	}
	
	// Extract keywords (split by common delimiters)
	words := strings.FieldsFunc(prompt, func(c rune) bool {
		return c == ' ' || c == '，' || c == '、' || c == '。' || c == '！' || c == '？'
	})
	
	for _, word := range words {
		word = strings.TrimSpace(word)
		if len(word) > 1 {
			analysis.Keywords = append(analysis.Keywords, word)
		}
	}
	
	// Detect type based on keywords
	if containsAny(prompt, []string{"猫", "狗", "鸟", "鱼", "动物", "cat", "dog", "bird", "animal"}) {
		analysis.Type = "animal"
	} else if containsAny(prompt, []string{"人", "男", "女", "孩子", "角色", "人物", "character", "person"}) {
		analysis.Type = "character"
	} else if containsAny(prompt, []string{"风景", "景色", "自然", "山", "海", "森林", "scene", "landscape", "nature"}) {
		analysis.Type = "scene"
	} else if containsAny(prompt, []string{"房子", "建筑", "城市", "楼", "building", "house", "city"}) {
		analysis.Type = "building"
	} else if containsAny(prompt, []string{"食物", "水果", "蔬菜", "食品", "food", "fruit", "vegetable"}) {
		analysis.Type = "food"
	} else if containsAny(prompt, []string{"车", "飞机", "船", "交通", "vehicle", "car", "plane", "transport"}) {
		analysis.Type = "vehicle"
	} else if containsAny(prompt, []string{"花", "树", "植物", "grass", "flower", "tree", "plant"}) {
		analysis.Type = "nature"
	} else if containsAny(prompt, []string{"抽象", "艺术", "创意", "abstract", "art", "creative"}) {
		analysis.Type = "abstract"
	}
	
	// Detect emotion
	if containsAny(prompt, []string{"可爱", "萌", "甜美", "cute", "sweet", "adorable"}) {
		analysis.Emotion = "cute"
	} else if containsAny(prompt, []string{"严肃", "正式", "专业", "serious", "formal", "professional"}) {
		analysis.Emotion = "serious"
	} else if containsAny(prompt, []string{"恐怖", "可怕", "黑暗", "scary", "dark", "horror"}) {
		analysis.Emotion = "scary"
	} else if containsAny(prompt, []string{"快乐", "开心", "欢乐", "happy", "joyful", "cheerful"}) {
		analysis.Emotion = "happy"
	} else if containsAny(prompt, []string{"神秘", "魔法", "奇幻", "mysterious", "magical", "mystical"}) {
		analysis.Emotion = "mysterious"
	}
	
	// Detect complexity based on prompt length and detail words
	if len(prompt) > 50 || containsAny(prompt, []string{"详细", "复杂", "精致", "detailed", "complex", "intricate"}) {
		analysis.Complexity = "complex"
	} else if len(prompt) > 20 || containsAny(prompt, []string{"中等", "适中", "medium", "moderate"}) {
		analysis.Complexity = "medium"
	}
	
	return analysis
}

// containsAny checks if the text contains any of the given keywords
func containsAny(text string, keywords []string) bool {
	for _, keyword := range keywords {
		if strings.Contains(text, keyword) {
			return true
		}
	}
	return false
}

// OptimizePromptWithAnalysis performs multi-layer prompt optimization based on analysis
func OptimizePromptWithAnalysis(userPrompt string, theme ThemeType) string {
	if theme == ThemeNone {
		return userPrompt
	}
	
	// Step 1: Analyze prompt
	analysis := AnalyzePromptType(userPrompt)
	
	// Step 2: Get theme enhancement
	themePrompt := GetThemePromptEnhancement(theme)
	
	// Step 3: Get quality enhancement based on complexity
	qualityPrompt := QualityPrompts[analysis.Complexity]
	
	// Step 4: Get style enhancement based on type
	stylePrompt := StylePrompts[analysis.Type]
	
	// Step 5: Get technical requirements
	technicalPrompt := TechnicalPrompts[theme]
	
	// Step 6: Combine all prompts intelligently
	var promptParts []string
	
	// Base prompt (user input)
	promptParts = append(promptParts, userPrompt)
	
	// Theme prompt (most important for theme consistency)
	if themePrompt != "" {
		promptParts = append(promptParts, themePrompt)
	}
	
	// Quality prompt
	if qualityPrompt != "" {
		promptParts = append(promptParts, qualityPrompt)
	}
	
	// Style prompt (emotion and type specific)
	if stylePrompt != "" {
		promptParts = append(promptParts, stylePrompt)
	}
	
	// Technical prompt
	if technicalPrompt != "" {
		promptParts = append(promptParts, technicalPrompt)
	}
	
	return strings.Join(promptParts, "，")
}

// GetOptimizedPromptForSearch returns an optimized prompt specifically for search
func GetOptimizedPromptForSearch(userPrompt string, theme ThemeType, searchType string) string {
	if searchType == "semantic" {
		// For semantic search, keep the prompt closer to original to maintain relevance
		if theme == ThemeNone {
			return userPrompt
		}
		// Light theme enhancement for semantic search
		themePrompt := GetThemePromptEnhancement(theme)
		return fmt.Sprintf("%s，%s", userPrompt, themePrompt)
	} else if searchType == "theme" {
		// For theme search, use full optimization to ensure style consistency
		return OptimizePromptWithAnalysis(userPrompt, theme)
	}
	
	return userPrompt
}
