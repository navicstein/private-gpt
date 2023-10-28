package helpers

import (
	"fmt"
	"testing"
)

func TestMDToHTML(t *testing.T) {
	var mds = `
	To complete the assignment, the following actions need to be taken:

1. Build three Virtual Machines (VM):
   - Linux Server: Use either CentOS or Ubuntu Server with minimum required configuration.
   - DNS: Configure and set up a DNS service.
   - Additional Service: Choose and configure an additional service of your choice, such as DHCP, FTP, SMTP, SNMP, etc.
   - Client: Create a client VM, which can be either Windows (Vista, 7, 8, 10, etc.) or Linux (Fedora, Ubuntu, etc.). Multiple copies of the client's VM can be created if more clients are needed to demonstrate an attack.
   - Attacker machine: Set up Kali Linux as the attacker machine (latest version is recommended).

2. Part A: Configuration and Attack Demonstration
   - Provide a summary of the configuration steps on the server and client.
   - Include screenshots to show the functionality at the client-side.
   - Discuss the rationale behind service selection and configuration (10% weightage).
   - Demonstrate a minimum of 2 attacks against each of the two services configured, with the option to perform more complex attacks for additional marks (35% weightage).
   - Log all important and offensive events against your target, including attacks detected, nature of service logs, origin of the attack, and damage caused. Support your demonstration with screenshots.
   - Critically reflect on the countermeasures and prevention mechanisms applied to mitigate against your attacks (15% weightage).

3. Part B: Position Paper on Recent Attacks and Countermeasures
   - Write a short position paper to critically analyze and reflect on recent state-of-the-art attacks and hacking techniques, followed by a discussion on possible countermeasures (40% weightage).
   - The paper should include a title page, abstract, introduction, scope, main body for critical discussion, reflection, and analysis, conclusions, and references. Referencing should follow the Harvard Style throughout the paper.

Overall, the actions involve building VMs, configuring services, demonstrating attacks, logging events, reflecting on countermeasures, and writing a position paper on recent attacks and countermeasures.
	`
	html := MDToHTML(mds)
	fmt.Println(html)
}
