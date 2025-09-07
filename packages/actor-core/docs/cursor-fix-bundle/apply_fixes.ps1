# apply_fixes.ps1 — optional helper for Cursor (best-effort search/replace)
# NOTE: This script is conservative; it will not overwrite if patterns aren't found.

Param(
  [string]$RepoRoot = "."
)

Function Replace-InFile {
  Param([string]$Path, [string]$Pattern, [string]$Replacement)
  if (Test-Path $Path) {
    $text = Get-Content -Raw -LiteralPath $Path
    if ($text -match [regex]::Escape($Pattern)) {
      $new = $text -replace [regex]::Escape($Pattern), [System.Text.RegularExpressions.Regex]::Escape($Replacement) -replace '\\\\', '\'
      if ($new -ne $text) {
        Set-Content -LiteralPath $Path -Encoding UTF8 -NoNewline -Value $new
        Write-Host "Patched: $Path"
      } else {
        Write-Host "No change (already applied): $Path"
      }
    } else {
      Write-Host "Pattern not found in $Path"
    }
  } else {
    Write-Host "File not found: $Path"
  }
}

# 1) Remove literal ellipses from known files (safe op)
$targets = @(
  Join-Path $RepoRoot "packages/actor-core/src/actorcore.go"),
  Join-Path $RepoRoot "packages/actor-core/tests/actorcore_test.go"
)
foreach ($f in $targets) {
  if (Test-Path $f) {
    (Get-Content -Raw $f).Replace("...", "") | Set-Content -Encoding UTF8 -NoNewline $f
    Write-Host "Removed ellipses in $f"
  }
}

Write-Host "Done. Review remaining manual patches per RUNBOOK."
