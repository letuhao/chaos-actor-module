# 12b — Search & Replace Plan for Stat Keys

> Mục tiêu: Loại bỏ toàn bộ string literal stat key và thay bằng `constants.Stat_*` (Go) hoặc `Stat.*` (TS).

## Regex gợi ý (Go)
- Tìm literal snake_case: `"(?:[a-z]+_)+[a-z]+"`
- Filter tiền xử lý: chỉ trong `services/`, `models/`, `interfaces/`, `constants/` (trừ file *_gen.go).

## Quy trình
1) Chạy codegen để tạo `constants/stats_gen.go` từ `tools/StatSchema.yml`.
2) Duyệt từng match, map thủ công hoặc bán tự động:
   - `"hp_max"` → `constants.Stat_HP_MAX`
   - `"move_speed"` → `constants.Stat_MOVE_SPEED`
3) Commit nhỏ theo nhóm stat (hp_*, atk_*, def_*, crit_*, speed_*, regen_*).

> Tip PowerShell (chỉ in kết quả, không sửa):  
```powershell
Get-ChildItem -Recurse -Include *.go | 
  Select-String -Pattern '"(?:[a-z]+_)+[a-z]+"' | 
  Where-Object { $_.Path -notmatch '_gen\.go$' } |
  Sort-Object Path, LineNumber |
  Format-Table -AutoSize
```
