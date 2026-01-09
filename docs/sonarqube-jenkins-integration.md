
# SonarQube + Jenkins Integration – End-to-End Guide

## 1. What is SonarQube (LTS)?

**SonarQube LTS** is a **central code quality server**.

It is responsible for:
- Storing source code analysis results
- Applying quality rules
- Evaluating Quality Gates
- Showing reports in a Web UI
- Sending analysis results back to CI tools

SonarQube **does NOT**:
- Run tests
- Build code
- Scan repositories by itself

Think of SonarQube as the **decision-maker**.

---

## 2. What is sonar-scanner-cli?

`sonar-scanner-cli` is a **client tool**.

It:
1. Reads source code from the Jenkins workspace
2. Reads configuration (`sonar-project.properties`)
3. Reads coverage reports (e.g. `coverage.out`)
4. Sends analysis data to the SonarQube server

It does **not** decide pass/fail.

Think of the scanner as a **courier**.

---

## 3. Role of Jenkins

Jenkins acts as the **orchestrator**.

Jenkins:
- Checks out the code
- Runs unit tests
- Generates coverage reports
- Executes sonar-scanner
- Waits for the Quality Gate result
- Decides whether the pipeline continues or stops

Jenkins enforces policy, SonarQube defines policy.

---

## 4. End-to-End Flow (What Happens in Our Pipeline)

```
Developer pushes code
        ↓
Jenkins pipeline starts
        ↓
Go unit tests run (coverage.out generated)
        ↓
sonar-scanner-cli runs in Jenkins
        ↓
Analysis data sent to SonarQube LTS server
        ↓
SonarQube processes analysis asynchronously
        ↓
Quality Gate is evaluated
        ↓
SonarQube sends result to Jenkins (Webhook / Polling)
        ↓
Jenkins waits using waitForQualityGate
        ↓
If Quality Gate = OK → pipeline continues
If Quality Gate = FAILED → pipeline stops
```

---

## 5. Quality Gate Explained

A **Quality Gate** is a set of conditions such as:
- No new critical bugs
- No new vulnerabilities
- Minimum coverage on new code
- Maintainability rating

Possible results:
- `OK` → Green (pipeline continues)
- `ERROR` → Red (pipeline fails)

Quality Gates are **policies**, not tests.

---

## 6. How Jenkins Enforces the Quality Gate

Jenkins uses this stage:

```groovy
stage('Quality Gate') {
    steps {
        timeout(time: 5, unit: 'MINUTES') {
            waitForQualityGate abortPipeline: true
        }
    }
}
```

What this does:
- Jenkins waits for SonarQube’s decision
- If status is not `OK`, Jenkins aborts the pipeline
- Prevents building or pushing bad code

This is **fail-fast enforcement**.

---

## 7. Webhook vs Polling

### Webhook (Recommended)
- SonarQube notifies Jenkins immediately
- Faster and scalable
- Required for production setups

Webhook URL:
```
http://<jenkins-host>:8080/sonarqube-webhook/
```

### Polling (Fallback)
- Jenkins periodically asks SonarQube for status
- Works for small setups
- Slower and not ideal

---

## 8. Why SonarQube Runs Before Build

Correct order:
```
Test → SonarQube → Quality Gate → Build → Push → Deploy
```

Reasons:
- SonarQube analyzes **source code**, not binaries
- Avoids building images for bad-quality code
- Saves time and resources
- Enforces governance

---

## 9. What We Implemented

In our setup, we:
- Ran Go unit tests with coverage
- Configured `sonar-project.properties`
- Used `sonar-scanner-cli` in Jenkins
- Connected Jenkins to SonarQube using a token
- Enforced Quality Gates using `waitForQualityGate`
- Ensured pipeline blocks on quality failure

This is a **production-grade CI integration**.

---

## 10. One-Line Summary (Interview Ready)

> Jenkins runs tests and sonar-scanner, SonarQube evaluates code quality using Quality Gates, and Jenkins enforces the result by blocking or allowing the pipeline.

---

## 11. Key Takeaway

**Scanner talks → Server decides → Jenkins enforces**

This separation makes the system scalable, reliable, and enterprise-ready.
