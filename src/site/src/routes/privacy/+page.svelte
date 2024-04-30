<svelte:head>
	<title>OTel Recipes - Privacy</title>
</svelte:head>

<section class="section">
	<div class="container">
		<div class="columns">
			<div class="column is-8">
				<h3 class="title is-3">Privacy policy</h3>

				<div class="block">
					This page explains what visiting this site means for your online privacy. First and
					foremost, online privacy is something <strong>very important to me</strong>. This site was
					built with your online privacy as a priority from the start. This site is entirely Open
					Source and its code can be found on <a href="https://github.com/joaopgrassi/otel-recipes">GitHub</a>.
				</div>

				<div class="block">
					Although I have <strong>no interest</strong> in collecting nor sharing any of your data,
					due to technicalities <em>some</em> data might be transmitted. The next sections describe in
					detail the scope of such data transmission.
				</div>

				<h1 class="title is-4">1. Data controller</h1>

				<div class="block">
					<p>Joao Grassi</p>

					<p>privacy (at) otelrecipes (dot) com</p>
				</div>

				<h1 class="title is-4">2. Accessing this site</h1>

				<h1 class="title is-5">2.1 Host provider</h1>

				<div class="block">
					Data processor: <a href="https://www.hetzner.com/legal/legal-notice"
						>Hetzner Online GmbH</a
					> with servers within Germany/Finland. When visiting the site, the following data is stored
					in the host provider's log files:
				</div>
				<div class="block content">
					<ul>
						<li>Anonymized IP address</li>
						<li>Browser</li>
						<li>Date and time of visit</li>
						<li>Operating system</li>
					</ul>
				</div>

				<div class="block content">
					From the provider's <a
						href="https://docs.hetzner.com/general/general-terms-and-conditions/data-privacy-faq/#which-data-is-stored-in-the-log-files"
						>FAQ:</a
					>
					<blockquote class="mt-5">
						At the web server level, this happens by default by storing an IP address
						123.123.123.XXX in the log file instead of the visitor's actual IP address, for example,
						123.123.123.123. The XXX is a random value between 1 and 254, so it is no longer
						possible to establish the true identity of the visitor.
					</blockquote>
				</div>

				<div class="block content">
					Purposes of processing:
					<ul>
						<li>Ensuring the function of the website</li>
						<li>Traceability and protection against attacks</li>
						<li>Legitimate interest based on Article 6(1)(f) GDPR</li>
					</ul>
					Data is stored for:<strong>14 days in encrypted backups**</strong>.
				</div>

				<h1 class="title is-5">2.2 Analytics</h1>

				<div class="block content">
					Data processor: <a href="https://www.goatcounter.com/help/privacy">GoatCounter</a> When
					visiting the site, the following data can be stored by the provider:
					<ul>
						<li>URL of the visited page</li>
						<li><code>Referer</code> header</li>
						<li>
							Browser and system information (derived from User-Agent header or HTTP client hints;
							the original headers are not stored)
						</li>
						<li>Screen size</li>
						<li>Country and region name derived from the IP address</li>
					</ul>
				</div>

				<div class="block content">
					From the provider's <a href="https://www.goatcounter.com/help/privacy">Privacy policy:</a>
					<blockquote class="mt-5">
						<strong>No personal information (such as IP address) is collected.</strong> A hash of
						the
						<code>IP address</code>, <code>User-Agent</code>, and a random number (“salt”) is kept
						in the process memory for 8 hours to identify a browsing session, and is never stored to
						disk.
					</blockquote>
				</div>

				<div class="block content">
					The usage of GoatCounter does not require a consent on the basis that
					<ol>
						<li>No personally identifiable information is collected</li>
						<li>
							It is in the legitimate interest of a site’s owner to see how many people are visiting
							their site
						</li>
					</ol>
					A more in depth explanation can be found GoutCounter's <a
						href="https://trinity.goatcounter.com/help/gdpr"
					>
						GDPR consent notices
					</a>
					page.
				</div>
			</div>
		</div>
	</div>
</section>
