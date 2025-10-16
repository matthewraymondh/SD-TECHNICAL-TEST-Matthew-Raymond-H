# Contributing to XYZ Football API

Terima kasih atas minat Anda untuk berkontribusi pada XYZ Football API! 🎉

## 🤝 Cara Berkontribusi

### 1. Fork Repository
Fork repository ini ke akun GitHub Anda.

### 2. Clone Repository
```bash
git clone https://github.com/your-username/xyz-football-api.git
cd xyz-football-api
```

### 3. Buat Branch Baru
```bash
git checkout -b feature/nama-fitur-anda
```

Naming convention untuk branch:
- `feature/` - untuk fitur baru
- `bugfix/` - untuk perbaikan bug
- `docs/` - untuk update dokumentasi
- `refactor/` - untuk refactoring code

### 4. Lakukan Perubahan
Pastikan code Anda mengikuti best practices:
- Code yang clean dan readable
- Tambahkan comments untuk code yang kompleks
- Follow Go conventions dan idioms
- Test code Anda sebelum commit

### 5. Commit Changes
```bash
git add .
git commit -m "feat: deskripsi singkat perubahan"
```

Commit message format:
- `feat:` - fitur baru
- `fix:` - bug fix
- `docs:` - perubahan dokumentasi
- `refactor:` - refactoring
- `test:` - menambah tests
- `chore:` - maintenance

### 6. Push ke GitHub
```bash
git push origin feature/nama-fitur-anda
```

### 7. Buat Pull Request
- Buka repository Anda di GitHub
- Klik "New Pull Request"
- Pilih branch Anda
- Isi deskripsi lengkap tentang perubahan
- Submit PR

## 📝 Guidelines

### Code Style
- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use `gofmt` untuk format code
- Use `golint` untuk linting
- Gunakan nama variable yang descriptive

### Testing
- Tambahkan unit tests untuk fitur baru
- Pastikan semua tests pass sebelum submit PR
- Test coverage minimal 70%

### Documentation
- Update README.md jika ada perubahan API
- Tambahkan GoDoc comments untuk public functions
- Update CHANGELOG.md

### Commit Messages
Gunakan format conventional commits:
```
<type>(<scope>): <subject>

<body>

<footer>
```

Contoh:
```
feat(auth): add refresh token endpoint

Implement refresh token functionality to allow users
to get new access token without re-login.

Closes #123
```

## 🐛 Melaporkan Bug

Jika Anda menemukan bug:

1. Cek [Issues](https://github.com/xyz/xyz-football-api/issues) apakah bug sudah dilaporkan
2. Jika belum, buat issue baru dengan template:

```markdown
**Deskripsi Bug:**
Penjelasan singkat tentang bug

**Cara Reproduksi:**
1. Step 1
2. Step 2
3. ...

**Expected Behavior:**
Apa yang seharusnya terjadi

**Actual Behavior:**
Apa yang terjadi

**Environment:**
- OS: 
- Go Version:
- Database Version:

**Screenshots:**
(jika ada)
```

## 💡 Mengusulkan Fitur

Untuk mengusulkan fitur baru:

1. Buat issue dengan label `enhancement`
2. Jelaskan use case dan benefit dari fitur
3. Diskusikan dengan maintainers sebelum implement

## 🔍 Code Review Process

Setelah submit PR:

1. Maintainer akan review code Anda
2. Mungkin ada request untuk perubahan
3. Update PR berdasarkan feedback
4. Setelah approved, PR akan di-merge

## ✅ Checklist Sebelum Submit PR

- [ ] Code sudah di-test dan berjalan dengan baik
- [ ] Code sudah di-format dengan `gofmt`
- [ ] Tidak ada error dari `golint`
- [ ] Dokumentasi sudah diupdate
- [ ] Commit messages mengikuti convention
- [ ] Branch sudah up-to-date dengan main
- [ ] Tests sudah ditambahkan (untuk fitur baru)

## 🎯 Area yang Perlu Kontribusi

Kami sangat appreciate kontribusi di area berikut:

- [ ] Unit tests & integration tests
- [ ] API documentation (Swagger/OpenAPI)
- [ ] Performance optimization
- [ ] Docker support
- [ ] CI/CD pipeline
- [ ] Additional features
- [ ] Bug fixes
- [ ] Documentation improvements

## 📞 Komunikasi

- **Issues**: Untuk bug reports dan feature requests
- **Discussions**: Untuk pertanyaan dan diskusi
- **Email**: [your-email@example.com] untuk urgent matters

## 📜 Code of Conduct

Dengan berkontribusi, Anda setuju untuk:

- Bersikap hormat kepada kontributor lain
- Memberikan feedback yang konstruktif
- Fokus pada apa yang terbaik untuk project
- Menerima kritik dengan gracefully

## 🙏 Terima Kasih!

Setiap kontribusi, sekecil apapun, sangat berarti bagi project ini!

---

**Happy Contributing!** 🚀⚽
