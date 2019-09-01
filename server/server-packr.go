package server

import "github.com/gobuffalo/packr"

// !!! GENERATED FILE !!!
// Do NOT hand edit this file!!
// It is recommended that you DO NOT check into this file into SCM.
// We STRONGLY recommend you delete this file after you have built your
// Go binary. You can use the "packr clean" command to clean up this,
// and any other packr generated files.
func init() {
	packr.PackJSONBytes("./static/templates", "error.html", "\"PGh0bWw+CjxoZWFkPgogIDx0aXRsZT5CYWQgR2F1Y2hlIExpbmtzPC90aXRsZT4KICA8bGluayByZWw9Imljb24iIHR5cGU9ImltYWdlL3BuZyIgaHJlZj0iL19mYXZpY29uLnBuZyI+CiAgPGxpbmsgcmVsPSJzdHlsZXNoZWV0IiBocmVmPSJodHRwczovL21heGNkbi5ib290c3RyYXBjZG4uY29tL2Jvb3RzdHJhcC80LjAuMC9jc3MvYm9vdHN0cmFwLm1pbi5jc3MiIGludGVncml0eT0ic2hhMzg0LUduNTM4NHhxUTFhb1dYQSswNThSWFB4UGc2Znk0SVd2VE5oMEUyNjNYbUZjSmxTQXdpR2dGQVcvZEFpUzZKWG0iIGNyb3Nzb3JpZ2luPSJhbm9ueW1vdXMiPgogIDxtZXRhIG5hbWU9InZpZXdwb3J0IiBjb250ZW50PSJ3aWR0aD1kZXZpY2Utd2lkdGgsIGluaXRpYWwtc2NhbGU9MSwgc2hyaW5rLXRvLWZpdD1ubyI+CjwvaGVhZD4KPGJvZHk+CjxkaXYgY2xhc3M9ImNvbnRhaW5lciI+CiAgPGRpdiBjbGFzcz0icm93Ij4KICAgIDxkaXYgY2xhc3M9ImFsZXJ0IGFsZXJ0LWRhbmdlciIgcm9sZT0iYWxlcnQiPgogICAgICBUaGVyZSB3YXMgYW4gZXJyb3IgcHJvY2Vzc2luZyB5b3VyIGxpbmsgJ3t7IC5saW5rLlBhdGggfX0nIHdpdGggdGFyZ2V0ICc8YSBocmVmPSJ7eyAubGluay5UYXJnZXQgfX0iPnt7LmxpbmsuVGFyZ2V0fX08L2E+Jzoge3sgLmVyciB9fS4gIDxhIGhyZWY9Int7IC5saW5rLkVkaXRVUkwgfX0iPkVkaXQgaXQ/PC9hPgogICAgPC9kaXY+CiAgPC9kaXY+CjwvZGl2Pgo8L2JvZHk+CjwvaHRtbD4=\"")
	packr.PackJSONBytes("./static/templates", "index.html", "\"PGh0bWw+CjxoZWFkIHByb2ZpbGU9Imh0dHA6Ly9hOS5jb20vLS9zcGVjL29wZW5zZWFyY2gvMS4xLyI+CiAgPHRpdGxlPkdhdWNoZUxpbmtzPC90aXRsZT4KICA8bGluayByZWw9Imljb24iIHR5cGU9ImltYWdlL3BuZyIgaHJlZj0iL19mYXZpY29uLnBuZyI+CiAgPGxpbmsgcmVsPSJzdHlsZXNoZWV0IiBocmVmPSJodHRwczovL21heGNkbi5ib290c3RyYXBjZG4uY29tL2Jvb3RzdHJhcC80LjAuMC9jc3MvYm9vdHN0cmFwLm1pbi5jc3MiCiAgICAgICAgaW50ZWdyaXR5PSJzaGEzODQtR241Mzg0eHFRMWFvV1hBKzA1OFJYUHhQZzZmeTRJV3ZUTmgwRTI2M1htRmNKbFNBd2lHZ0ZBVy9kQWlTNkpYbSIgY3Jvc3NvcmlnaW49ImFub255bW91cyI+CiAgPG1ldGEgbmFtZT0idmlld3BvcnQiIGNvbnRlbnQ9IndpZHRoPWRldmljZS13aWR0aCwgaW5pdGlhbC1zY2FsZT0xLCBzaHJpbmstdG8tZml0PW5vIj4KICA8bGluayByZWw9InNlYXJjaCIgdHlwZT0iYXBwbGljYXRpb24vb3BlbnNlYXJjaGRlc2NyaXB0aW9uK3htbCIgaHJlZj0ie3sgLmhvc3QgfX0vX3NlYXJjaC54bWwiIHRpdGxlPSJDb250ZW50IHNlYXJjaCIvPgogIDxzY3JpcHQgZGVmZXIgc3JjPSJodHRwczovL3VzZS5mb250YXdlc29tZS5jb20vcmVsZWFzZXMvdjUuMC42L2pzL2FsbC5qcyI+PC9zY3JpcHQ+CiAgPHNjcmlwdCBzcmM9Imh0dHBzOi8vY29kZS5qcXVlcnkuY29tL2pxdWVyeS0zLjIuMS5zbGltLm1pbi5qcyIKICAgICAgICAgIGludGVncml0eT0ic2hhMzg0LUtKM28yREt0SWt2WUlLM1VFTnptTTdLQ2tSci9yRTkvUXBnNmFBWkdKd0ZETVZOQS9HcEdGRjkzaFhwRzVLa04iCiAgICAgICAgICBjcm9zc29yaWdpbj0iYW5vbnltb3VzIj48L3NjcmlwdD4KICA8c2NyaXB0IHNyYz0iaHR0cHM6Ly9jZG5qcy5jbG91ZGZsYXJlLmNvbS9hamF4L2xpYnMvcG9wcGVyLmpzLzEuMTIuOS91bWQvcG9wcGVyLm1pbi5qcyIKICAgICAgICAgIGludGVncml0eT0ic2hhMzg0LUFwTmJnaDlCK1kxUUt0djNSbjdXM21nUHhoVTlLL1NjUXNBUDdoVWliWDM5ajdmYWtGUHNrdlh1c3ZmYTBiNFEiCiAgICAgICAgICBjcm9zc29yaWdpbj0iYW5vbnltb3VzIj48L3NjcmlwdD4KICA8c2NyaXB0IHNyYz0iaHR0cHM6Ly9tYXhjZG4uYm9vdHN0cmFwY2RuLmNvbS9ib290c3RyYXAvNC4wLjAvanMvYm9vdHN0cmFwLm1pbi5qcyIKICAgICAgICAgIGludGVncml0eT0ic2hhMzg0LUpaUjZTcGVqaDRVMDJkOGpPdDZ2TEVIZmUvSlFHaVJSU1FReFNmRldwaTFNcXVWZEF5alVhcjUrNzZQVkNtWWwiCiAgICAgICAgICBjcm9zc29yaWdpbj0iYW5vbnltb3VzIj48L3NjcmlwdD4KPC9oZWFkPgo8Ym9keT4KPHNjcmlwdCB0eXBlPSJhcHBsaWNhdGlvbi9sZCtqc29uIj4KPC9zY3JpcHQ+CjxkaXYgY2xhc3M9ImNvbnRhaW5lciI+CiAgPGRpdiBjbGFzcz0icm93Ij4KICAgIDxkaXYgY2xhc3M9ImNvbC1zbS0xMiI+CiAgICAgIDxwPjwvcD4KICAgICAgPGgxIGNsYXNzPSJ0ZXh0LWNlbnRlciB0ZXh0LXByaW1hcnkiPgogICAgICAgIDxpIGNsYXNzPSJmYXMgZmEtcm9ja2V0IGZhLWZsaXAtaG9yaXpvbnRhbCI+PC9pPgogICAgICAgIDxpIGNsYXNzPSJmYXMgZmEtcm9ja2V0IGZhLWZsaXAtaG9yaXpvbnRhbCI+PC9pPgogICAgICAgIDxpIGNsYXNzPSJmYXMgZmEtcm9ja2V0IGZhLWZsaXAtaG9yaXpvbnRhbCI+PC9pPgogICAgICAgIEdhdWNoZUxpbmtzCiAgICAgICAgPGkgY2xhc3M9ImZhcyBmYS1yb2NrZXQgZmEtZmxpcC1ob3Jpem9udGFsIj48L2k+CiAgICAgICAgPGkgY2xhc3M9ImZhcyBmYS1yb2NrZXQgZmEtZmxpcC1ob3Jpem9udGFsIj48L2k+CiAgICAgICAgPGkgY2xhc3M9ImZhcyBmYS1yb2NrZXQgZmEtZmxpcC1ob3Jpem9udGFsIj48L2k+CgogICAgICA8L2gxPgoKICAgIHt7IGlmIGVxIC5xdWVyeSAiIiB9fQogICAgICA8YSBjbGFzcz0iYnRuIGJ0bi1vdXRsaW5lLXByaW1hcnkiIHJvbGU9ImJ1dHRvbiIKICAgICAgICAgaHJlZj0iL19hZGQ/cHJlZmlsbF9QYXRoPXt7IC5xdWVyeSB9fSZwcmVmaWxsX0F1dGhvcj17eyAuYXV0aG9yIH19Ij5BZGQgYSBsaW5rPC9hPgogICAge3sgZWxzZSB9fQogICAgICA8YSBjbGFzcz0iYnRuIGJ0bi1vdXRsaW5lLXByaW1hcnkiIHJvbGU9ImJ1dHRvbiIKICAgICAgICAgaHJlZj0iL19hZGQ/cHJlZmlsbF9QYXRoPXt7IC5xdWVyeSB9fSZwcmVmaWxsX0F1dGhvcj17eyAuYXV0aG9yIH19Ij5BZGQgPHN0cm9uZz57eyAucHJlZml4IH19L3t7IC5xdWVyeSB9fTwvc3Ryb25nPiBhcyBhIG5ldyBsaW5rPC9hPgogICAge3sgZW5kIH19CgogICAge3sgJGRvd25sb2FkX2Nocm9tZV9tc2cgOj0gYAogICAgPG9sPgogICAgICA8bGk+RG93bmxvYWQgdGhlIGV4dGVuc2lvbiBmb3IgPGEgaHJlZj0iL19jaHJvbWVfZXh0ZW5zaW9uIj5DaHJvbWU8L2E+IG9yIDxhIGhyZWY9Ii9fZmlyZWZveF9leHRlbnNpb24iPkZpcmVmb3g8L2E+LgogICAgICA8bGk+T3BlbiA8YSBocmVmPSJjaHJvbWU6Ly9leHRlbnNpb25zIj5jaHJvbWU6Ly9leHRlbnNpb25zPC9hPiBieSBwYXN0aW5nIDxjb2RlPmNocm9tZTovL2V4dGVuc2lvbnM8L2NvZGU+IGludG8geW91ciBzZWFyY2ggYmFyLgogICAgICA8bGk+RW5hYmxlIGRldmVsb3BlciBtb2RlIHNvIHlvdSBjYW4gaW5zdGFsbCBhbiB1bnNpZ25lZCBleHRlbnNpb24uCiAgICAgIDxsaT5EcmFnIHRoZSBkb3dubG9hZGVkIGV4dGVuc2lvbiBpbnRvIHRoZSBleHRlbnNpb25zIHdpbmRvdy4KICAgIDwvb2w+CiAgICBgfX0KCiAgICAgIDxhIGhyZWY9Ii9fZWRpdCIgY2xhc3M9ImJ0biBidG4tbGluayI+RWRpdCBhbGwgbGlua3M8L2E+CiAgICAgIDxkaXYgY2xhc3M9ImZsb2F0LXJpZ2h0Ij4KICAgICAgICA8YSBjbGFzcz0iYnRuIGJ0bi1saW5rIGJ0bi1pbmZvIHRleHQtaW5mbyIgZGF0YS10b2dnbGU9InBvcG92ZXIiIHRpdGxlPSJIb3cgdG8gaW5zdGFsbCB0aGUgYnJvd3NlciBleHRlbnNpb24iCiAgICAgICAgICAgICAgICBkYXRhLWNvbnRlbnQ9Int7ICRkb3dubG9hZF9jaHJvbWVfbXNnIH19IgogICAgICAgID5Eb3dubG9hZCBCcm93c2VyIEV4dGVuc2lvbgogICAgICAgIDwvYT4KICAgICAgPC9kaXY+CiAgICAgIDxwPjwvcD4KICAgIHt7IGlmIC5saW5rcyB9fQogICAge3sgaWYgbmUgLnF1ZXJ5ICIiIH19CiAgICAgIE1heWJlIHlvdSBtZWFudCBvbmUgb2YgdGhlc2U/CiAgICB7eyBlbmQgfX0KICAgICAgPGRpdj4KICAgICAgICA8dGFibGUgY2xhc3M9InRhYmxlIHRhYmxlLXN0cmlwZWQgdGFibGUtcmVzcG9uc2l2ZSB0YWJsZS1jb25kZW5zZWQgdGFibGUtaG92ZXIiPgogICAgICAgICAgPHRoZWFkPgogICAgICAgICAgPHRyPgogICAgICAgICAgICA8dGg+TGluazwvdGg+CiAgICAgICAgICAgIDx0aD5EZXNjcmlwdGlvbjwvdGg+CiAgICAgICAgICAgIDx0aD5UYXJnZXQ8L3RoPgogICAgICAgICAgICA8dGg+PC90aD4KICAgICAgICAgIDwvdHI+CiAgICAgICAgICA8L3RoZWFkPgogICAgICAgICAgPHRib2R5PgogICAgICAgICAge3sgJG9yaWdpbmFsVVJMIDo9IC5vcmlnaW5hbFVSTCB9fQogICAgICAgICAge3sgJHByZWZpeCA6PSAucHJlZml4fX17eyByYW5nZSAkbCA6PSAubGlua3MgfX0KICAgICAgICAgIHt7ICRiIDo9ICRsLkJhc2UgfX0KICAgICAgICAgIDx0cj4KICAgICAgICAgICAgPHRkPjxhIGhyZWY9Int7ICRsLlRyYW5zZm9ybSAiIiAkb3JpZ2luYWxVUkwgfX0iPnt7ICRwcmVmaXggfX0ve3sgaWYgJGwuSXNSZWdleHAgfX0KICAgICAgICAgICAgICA8Y29kZSBjbGFzcz0iZm9udC1pdGFsaWMgdGV4dC1zdWNjZXNzIj57eyRsLlBhdGhTdHJpbmd9fTwvY29kZT57eyBlbHNlIH19e3sgJGIuUGF0aCB9fXt7IGVuZCB9fTwvYT48L3RkPgogICAgICAgICAgICA8dGQ+e3sgJGIuRGVzY3JpcHRpb24gfX08L3RkPgogICAgICAgICAgICA8dGQ+PGEgaHJlZj0ie3sgJGwuVHJhbnNmb3JtICIiICRvcmlnaW5hbFVSTCB9fSI+e3sgJGIuVGFyZ2V0IH19PC9hPjwvdGQ+CiAgICAgICAgICAgIDx0ZD48YSBocmVmPSJ7eyAkYi5FZGl0VVJMLlN0cmluZyB9fSI+PGkgY2xhc3M9ImZhciBmYS1zbSBmYS1lZGl0Ij48L2k+PC9hPjwvdGQ+CiAgICAgICAgICA8L3RyPgogICAgICAgICAge3sgZW5kIH19CiAgICAgICAgICA8L3Rib2R5PgogICAgICAgIDwvdGFibGU+CiAgICAgIDwvZGl2PgogICAge3sgZWxzZSB9fQogICAgICA8ZGl2IGNsYXNzPSJjYXJkIj4KICAgICAgICA8ZGl2IGNsYXNzPSJjYXJkLWJvZHkgdGV4dC1jZW50ZXIiPgogICAgICAgICAgTm8gbWF0Y2hlcyBmb3VuZCBmb3IKICAgICAgICAgIDxtYXJrPnt7IC5xdWVyeSB9fTwvbWFyaz4KICAgICAgICAgIC4KICAgICAgICA8L2Rpdj4KICAgICAgPC9kaXY+CiAgICB7eyBlbmQgfX0KICAgIDwvZGl2PgogIDwvZGl2Pgo8L2Rpdj4KPHNjcmlwdD4KICAkKGZ1bmN0aW9uICgpIHsKICAgICQoJ1tkYXRhLXRvZ2dsZT0icG9wb3ZlciJdJykucG9wb3Zlcih7aHRtbDogdHJ1ZX0pCiAgfSkKPC9zY3JpcHQ+CjwvYm9keT4KPC9odG1sPgo=\"")
	packr.PackJSONBytes("./static/templates", "search.xml", "\"PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz4KPE9wZW5TZWFyY2hEZXNjcmlwdGlvbiB4bWxucz0iaHR0cDovL2E5LmNvbS8tL3NwZWMvb3BlbnNlYXJjaC8xLjEvIj4KICAgIDxTaG9ydE5hbWU+TEQgR2F1Y2hlTGlua3M8L1Nob3J0TmFtZT4KICAgIDxEZXNjcmlwdGlvbj5MYXVuY2hEYXJrbHkgR2F1Y2hlTGlua3M8L0Rlc2NyaXB0aW9uPgogICAgPElucHV0RW5jb2Rpbmc+VVRGLTg8L0lucHV0RW5jb2Rpbmc+CiAgICA8SW1hZ2Ugd2lkdGg9IjE2IiBoZWlnaHQ9IjE2IiB0eXBlPSJpbWFnZS94LWljb24iPnt7IC5ob3N0IH19L19mYXZpY29uLnBuZzwvSW1hZ2U+CiAgICA8VXJsIHR5cGU9ImFwcGxpY2F0aW9uL3Jzcyt4bWwiIG1ldGhvZD0iR0VUIiB0ZW1wbGF0ZT0iL3tzZWFyY2hUZXJtc30iLz4KICAgIDxVcmwgdHlwZT0iYXBwbGljYXRpb24veC1vcGVuc2VhcmNoK3htbCIgbWV0aG9kPSJHRVQiIHRlbXBsYXRlPSIve3NlYXJjaFRlcm1zfSIvPgo8L09wZW5TZWFyY2hEZXNjcmlwdGlvbj4K\"")
	}
