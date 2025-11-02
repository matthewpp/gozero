Prepare

🧩 Step-by-Step Installation For Window
1. ✅ Make sure winget is available

Open PowerShell (as Administrator) and run:

`winget --version`


If it shows a version (like v1.9.0), you’re good.
If not, update Windows or install App Installer from the Microsoft Store.
---

2. 💻 Search for Go

You can see what’s available:

`winget search Go`


You should see something like:

| Name       | Id               | Version | Source |
|------------|------------------|---------|--------|
| Go         | GoLang.Go        | x.y.z   | winget |

---

3. 📦 Install Go

Run:

`winget install --id GoLang.Go -e`


Explanation:

- `--id GoLang.Go` → installs the official Go package

- `-e` → ensures exact match (avoid similar names)

This will automatically:

- Download the latest Go release

- Add `C:\Program Files\Go\bin` to your PATH

---

4. 🧠 Verify the installation

After it finishes, restart PowerShell (or open a new one) and run:

`go version`


Expected output example:

`go version go1.23.2 windows/amd64`

---

5. 🚀 Set Path

`setx PATH "$env:PATH;C:\Program Files\Go\bin"
`

run go version again to confirm.

----

