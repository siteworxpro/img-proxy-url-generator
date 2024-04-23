package generator

func (g *Generator) generatePlainUrl(file string) (string, error) {
	return "plain/" + file, nil
}
