package server

import (
	"github.com/labstack/echo/v4"
	"github.com/swaggo/swag"

	"log/slog"
	"net/http"

	_ "github.com/ugabiga/swan/bootstrap/docs"
)

type OpenAPIHandler struct {
	logger *slog.Logger
}

func NewOpenAPIHandler(logger *slog.Logger) (*OpenAPIHandler, error) {
	return &OpenAPIHandler{
		logger: logger,
	}, nil
}

func (h OpenAPIHandler) SetRouters(router *echo.Group) {
	router.GET("/openapi/doc.json", h.OpenAPIJSONHandler)
	router.GET("/openapi", h.OpenAPIDocHandler)
}

func (h OpenAPIHandler) OpenAPIJSONHandler(c echo.Context) error {
	doc, err := swag.ReadDoc()
	if err != nil {
		return err
	}

	_, _ = c.Response().Writer.Write([]byte(doc))

	return nil
}

func (h OpenAPIHandler) OpenAPIDocHandler(c echo.Context) error {
	serverAddress := c.Scheme() + "://" + c.Request().Host
	specURL := serverAddress + "/openapi/doc.json"

	//<body style="height: 100vh; background-color=black">
	htmlContent := `<!doctype html>
<html lang="en" data-theme="dark">
<head>
	<meta charset="utf-8" />
	<meta name="referrer" content="same-origin" />
	<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
	<link rel="icon" type="image/svg+xml" href="https://go-fuego.github.io/fuego/img/logo.svg">
	<title>OpenAPI specification</title>
	<script src="https://unpkg.com/@stoplight/elements/web-components.min.js"></script>
	<link rel="stylesheet" href="https://unpkg.com/@stoplight/elements/styles.min.css" />
</head>
<body style="background-color: black; height: 100vh;">
	<elements-api
		apiDescriptionUrl="` + specURL + `"
		layout="responsive"
		router="hash"
		logo="https://go-fuego.github.io/fuego/img/logo.svg"
		tryItCredentialsPolicy="same-origin"
	/>

	<!-- Script to override inline styles -->
	<script>
		// Function to update the styles of elements
		function updateTokenStyles() {
			const tokenStrings = document.querySelectorAll('span.token.string');
			const tokenProperties = document.querySelectorAll('span.token.property');
			const tokenNumber = document.querySelectorAll('span.token.number');
			const tokenOperator = document.querySelectorAll('span.token.operator');

			if (tokenOperator.length > 0) {
				tokenOperator.forEach(function(el) {
					el.style.color = '#ffbfc5';
				});
			}

			if (tokenStrings.length > 0) {
				tokenStrings.forEach(function(el) {
					el.style.color = '#cfeeff';
				});
			}

			if (tokenNumber.length > 0) {
				tokenNumber.forEach(function(el) {
					el.style.color = '#cfeeff';
				});
			}

			// Update color for token properties
			if (tokenProperties.length > 0) {
				tokenProperties.forEach(function(el) {
					el.style.color = '#738bd6';
				});
			}

			// Retry if elements haven't rendered yet
			if (tokenStrings.length === 0 || tokenProperties.length === 0) {
				setTimeout(updateTokenStyles, 50);
			}
		}

		// Listen for URL hash changes and reapply styles
		window.addEventListener("hashchange", function() {
			// Re-run the style update function when the URL changes
			updateTokenStyles();
		});

		// You can also detect navigation or route changes (like in SPAs) by observing DOM mutations
		const observer = new MutationObserver(function() {
			updateTokenStyles();
		});

		// Start observing changes in the document body
		observer.observe(document.body, { childList: true, subtree: true });
	</script>
</body>
</html>`
	return c.HTMLBlob(http.StatusOK, []byte(htmlContent))
}
