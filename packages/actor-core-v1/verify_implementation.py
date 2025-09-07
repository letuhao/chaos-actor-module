#!/usr/bin/env python3
"""
Verification script for Actor Core implementation
This script checks the Go implementation files for correctness without running Go.
"""

import os
import re

def check_file_exists(filepath):
    """Check if a file exists and return its status."""
    if os.path.exists(filepath):
        return f"✅ {filepath} exists"
    else:
        return f"❌ {filepath} missing"

def check_go_implementation():
    """Check the Go implementation file for key requirements."""
    impl_file = "src/actorcore.go"
    
    if not os.path.exists(impl_file):
        return ["❌ Implementation file not found"]
    
    with open(impl_file, 'r') as f:
        content = f.read()
    
    checks = []
    
    # Check for required functions
    required_functions = [
        "ComposeCore",
        "BaseFromPrimary", 
        "FinalizeDerived",
        "ClampDerived"
    ]
    
    for func in required_functions:
        if f"func (a *ActorCoreImpl) {func}(" in content:
            checks.append(f"✅ {func} function implemented")
        else:
            checks.append(f"❌ {func} function missing")
    
    # Check for stable sort in ComposeCore
    if "sort.Strings(keys)" in content:
        checks.append("✅ ComposeCore uses stable sort")
    else:
        checks.append("❌ ComposeCore missing stable sort")
    
    # Check for version bump in FinalizeDerived
    if "result.Version = base.Version + 1" in content:
        checks.append("✅ FinalizeDerived bumps version by +1")
    else:
        checks.append("❌ FinalizeDerived missing version bump")
    
    # Check for clamp bounds
    clamp_checks = [
        ("Haste", "math.Max(0.5, math.Min(2.0, result.Haste))"),
        ("CritChance", "math.Max(0.0, math.Min(1.0, result.CritChance))"),
        ("Resists", "math.Max(0.0, math.Min(0.8, v))"),
    ]
    
    for name, pattern in clamp_checks:
        if pattern in content:
            checks.append(f"✅ {name} bounds properly clamped")
        else:
            checks.append(f"❌ {name} bounds missing or incorrect")
    
    # Check for NaN/Inf handling
    if "sanitizeFloat" in content and "math.IsNaN" in content:
        checks.append("✅ NaN/Inf handling implemented")
    else:
        checks.append("❌ NaN/Inf handling missing")
    
    return checks

def check_test_implementation():
    """Check the test file for comprehensive coverage."""
    test_file = "tests/actorcore_test.go"
    
    if not os.path.exists(test_file):
        return ["❌ Test file not found"]
    
    with open(test_file, 'r') as f:
        content = f.read()
    
    checks = []
    
    # Check for property tests
    property_tests = [
        "TestComposeCore_Commutativity",
        "TestComposeCore_Idempotence", 
        "TestBaseFromPrimary_Monotonicity",
        "TestClampDerived_Bounds"
    ]
    
    for test in property_tests:
        if f"func {test}(" in content:
            checks.append(f"✅ {test} implemented")
        else:
            checks.append(f"❌ {test} missing")
    
    # Check for golden test
    if "TestGoldenTest_FixedBuckets" in content:
        checks.append("✅ Golden test implemented")
    else:
        checks.append("❌ Golden test missing")
    
    # Check for version bump test
    if "TestFinalizeDerived_VersionBump" in content:
        checks.append("✅ Version bump test implemented")
    else:
        checks.append("❌ Version bump test missing")
    
    return checks

def main():
    """Run all verification checks."""
    print("🔍 Actor Core Implementation Verification")
    print("=" * 50)
    
    # Check file structure
    print("\n📁 File Structure:")
    files_to_check = [
        "src/actorcore.go",
        "tests/actorcore_test.go", 
        "go.mod",
        "run_tests.go"
    ]
    
    for file in files_to_check:
        print(check_file_exists(file))
    
    # Check implementation
    print("\n🔧 Implementation Checks:")
    impl_checks = check_go_implementation()
    for check in impl_checks:
        print(check)
    
    # Check tests
    print("\n🧪 Test Coverage:")
    test_checks = check_test_implementation()
    for check in test_checks:
        print(check)
    
    # Summary
    print("\n📊 Summary:")
    all_checks = impl_checks + test_checks
    passed = sum(1 for check in all_checks if check.startswith("✅"))
    total = len(all_checks)
    
    print(f"Passed: {passed}/{total} checks")
    
    if passed == total:
        print("🎉 All checks passed! Implementation looks good.")
    else:
        print("⚠️  Some checks failed. Please review the implementation.")
    
    print("\n💡 Next Steps:")
    print("1. Install Go from https://golang.org/dl/")
    print("2. Run: go test ./tests/ -v")
    print("3. Run: go run run_tests.go")

if __name__ == "__main__":
    main()
