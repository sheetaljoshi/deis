:title: Planning Process
:description: The open source planning process for the Deis project used to define the roadmap and its execution.

.. _planning:

Planning Process
================
Deis features a lightweight process that emphasizes openness and ensures every community member can be an integral part of planning for the future.

The Role of Maintainers
-----------------------
`Maintainers`_ lead the project. Their duties include proposing the Roadmap, reviewing and integrating contributions and maintaining the vision of the project.

Open Roadmap
------------
The :ref:`roadmap` is a community document. While Maintainers propose the Roadmap, it gets discussed and refined in Release Planning Meetings.

Contributing to the Roadmap
---------------------------
Proposals and issues can be opened by anyone. Every member of the :ref:`community` is welcome to participate in the discussion by providing feedback and/or offering counter-proposals.

Release Milestones
------------------
The Roadmap gets delivered progressively via the :ref:`release_schedule`.  Releases are defined during Release Planning Meetings and managed using GitHub Milestones which track specific deliverables and work-in-progress.

Release Planning Meetings
-------------------------
Major decisions affecting the Roadmap are discussed during Release Planning Meetings on the first Thursday of each month, aligned with the :ref:`release_schedule`.

Release Planning Meetings are open to the public with access coordinated via the #deis IRC channel on Freenode.
Notes from past meetings are below, along with links to a recording of the entire meeting on YouTube.

November 2016
~~~~~~~~~~~~~

The next public release planning meeting for Deis will take place on
**Thursday, November 3rd** at **12pm MDT/1900 UTC**. The `event`_ can be added
directly to your calendar.

October 2016
~~~~~~~~~~~~

- Introduction - Seth Goings
- Release Review - Jason Hansen
- Deis CI Overview - Jonathan Chauncey

Archive: https://www.youtube.com/watch?v=LhglIIh9izw

September 2016
~~~~~~~~~~~~~~

- Introduction - Gabe Monroy
- Meet Matt Tucker - Seth Goings & Matt Tucker
- Workflow 2.4 Release Review - Jason Hansen
- Autoscaling Preview - Helgi Þorbjörnsson
- Roadmap Process - Jason Hansen

Archive: https://www.youtube.com/watch?v=sDaAZGDcRgU

August 2016
~~~~~~~~~~~

- Introduction - Jason Hansen
- Megaboom Demo - Aaron Schlesinger
- Deployments Demo - Helgi Þorbjörnsson
- Workflow 2.3 Release Review - Jason Hansen
- Workflow 2.4 Up Next - Jason Hansen

Archive: https://www.youtube.com/watch?v=ZWkDpi76H-E

July 2016
~~~~~~~~~

- Workflow 2.1 and 2.2 Release Review
- Helm Alpha.2 and Alpha.3 Review

Archive: https://www.youtube.com/watch?v=BVHT03uQ1WU

June 2016
~~~~~~~~~

- Introduction & Community Updates - Gabe Monroy
- Helm Alpha 1 Demo - Michelle Noorali
- How Mozilla uses Deis v1 - Josh Mize
- Deis V1 Maintenance Update - Matt Boersma
- Deis Workflow Release Update - Jason Hansen

Archive: https://youtu.be/MtGSwsRnpcM

May 2016
~~~~~~~~

- Introduction & Community Updates
- LTS Update
- Beta 3 Release and Beta 4 Status
- Helm Update, Helm Classic and Kubernetes Helm
- Wrap-up!

Archive: https://www.youtube.com/watch?v=Jb-X_yVE2-w

April 2016 (Deis LTS Release)
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

- Intro
- Move to Slack https://slack.deis.io
- Notes on Deis Workflow Beta
- Road to Workflow Stable
- Deis Workflow e2e Testing and Improvements
- LTS Release (v1.13.0) for Deis v1
- Call for Community Demos!

Archive: https://www.youtube.com/watch?v=72g9PxiR0iU

March 2016 (Deis v2 Beta)
~~~~~~~~~~~~~~~~~~~~~~~~~

We demoed Deis v2 beta and shared the architecture of the Deis v2 CI/CD pipeline.

Archive: https://www.youtube.com/watch?v=rIF3v1MZkJg

February 2016
~~~~~~~~~~~~~

There was no public release planning meeting for February.

January 2016 (Deis 1.x LTS release and Deis v2 Alpha)
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

We're still taking feedback and suggestions on what our users would like to
see in the 1.x LTS release. Feel free to provide your input on the
`LTS release issue`_ on GitHub.

Over the last month we've been hard at work on Deis v2 which puts the Deis
workflow people have come to know and love on top of the `Kubernetes`_ platform.
Aaron demoed what we've done so far with the release of Deis v2 Alpha.
(`v2 alpha walkthrough cheatsheet`_)

If you've tried out v2 alpha, let us know what you think via the
`v2 alpha feedback`_ GitHub issue!

You can also take a look at our `v2 beta milestone`_ goals and chime in there.

Archive: https://www.youtube.com/watch?v=8LNVluUFh1M

December 2015 (Deis 1.x LTS release)
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

While much of our engineering efforts have shifted to complete Deis v2,
we plan to ship one final release of the 1.x branch. This release will be
a long-term support (LTS) release, receiving bug fixes and security updates
for the foreseeable future.

We invite our community to help us plan this release, and look forward to your
comments in the `LTS release issue`_ on GitHub.

Archive: https://www.youtube.com/watch?v=U70FOwJuIH4

November 2015 (Deis v1.12.2)
~~~~~~~~~~~~~~~~~~~~~~~~~~~~

No archive due to technical difficulties.

October 2015 (Deis v1.12)
~~~~~~~~~~~~~~~~~~~~~~~~~

Roadmap items completed: Stateless Logger

Roadmap items added: Deis v2

Archive: https://www.youtube.com/watch?v=GK1TzeQxGaY

September 2015 (Deis v1.11)
~~~~~~~~~~~~~~~~~~~~~~~~~~~

Roadmap items completed: Production Hardening

Roadmap items added: Etcd Hardening, Rigger

Archive: https://www.youtube.com/watch?v=oqTUDBmriDA

August 2015 (Deis v1.10)
~~~~~~~~~~~~~~~~~~~~~~~~

Roadmap items completed: Scheduling and Orchestration, Etcd 2, Networking v2

Roadmap items added: Internal Service Discovery, Permissions and Teams, New Default Scheduler

Archive: https://www.youtube.com/watch?v=a6tOrv4Uzz4

July 2015 (Deis v1.9)
~~~~~~~~~~~~~~~~~~~~~

Roadmap items completed: Pluggable Storage Subsystem, User-defined Health Checks

Roadmap items added: Stateless Logger, Production Hardening

Archive: https://www.youtube.com/watch?v=f_rJFWSFY5I

June 2015 (Deis v1.8)
~~~~~~~~~~~~~~~~~~~~~

Roadmap items added: Pluggable Storage Subsystem, Networking V2, Etcd 2, User-defined Health Checks

Archive: https://www.youtube.com/watch?v=nC-DyN1_II4

Credits
-------
Thanks to `Amy Lindburg`_ and our friends at `Docker`_ for their inspiration.

.. _`Amy Lindburg`: https://twitter.com/amylindburg
.. _`Docker`: https://www.docker.com/
.. _`event`: https://goo.gl/FGNnOC
.. _`LTS release issue`: https://github.com/deis/deis/issues/4776
.. _`Maintainers`: https://github.com/deis/deis/blob/master/MAINTAINERS.md
.. _`Kubernetes`: http://kubernetes.io/
.. _`v2 alpha feedback`: https://github.com/deis/deis/issues/4827
.. _`v2 alpha walkthrough cheatsheet`: https://gist.github.com/arschles/5b7a75a50938913d3eb1
.. _`v2 beta milestone`: https://github.com/deis/deis/issues/4809
