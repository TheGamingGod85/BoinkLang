# Software Requirements Specification (SRS)
## Indian Sign Language Translator App

---

### Revision History
| Date       | Version | Description          | Author                |
|------------|---------|----------------------|-----------------------|
| 2025-04-27 | 1.0     | Initial draft        | Aayushya Lakkadwala   |

---

## Table of Contents
1. [Introduction](#1-introduction)  
   1. [Purpose](#11-purpose)  
   2. [Scope](#12-scope)  
   3. [Definitions, Acronyms, and Abbreviations](#13-definitions-acronyms-and-abbreviations)  
   4. [References](#14-references)  
   5. [Overview](#15-overview)  
2. [Overall Description](#2-overall-description)  
   1. [Product Perspective](#21-product-perspective)  
   2. [Product Functions](#22-product-functions)  
   3. [User Classes and Characteristics](#23-user-classes-and-characteristics)  
   4. [Operating Environment](#24-operating-environment)  
   5. [Design and Implementation Constraints](#25-design-and-implementation-constraints)  
   6. [Assumptions and Dependencies](#26-assumptions-and-dependencies)  
3. [Specific Requirements](#3-specific-requirements)  
   1. [External Interface Requirements](#31-external-interface-requirements)  
       - [User Interfaces](#311-user-interfaces)  
       - [Hardware Interfaces](#312-hardware-interfaces)  
       - [Software Interfaces](#313-software-interfaces)  
       - [Communication Interfaces](#314-communication-interfaces)  
   2. [Functional Requirements](#32-functional-requirements)  
   3. [Non-functional Requirements](#33-non-functional-requirements)  
       - [Performance Requirements](#331-performance-requirements)  
       - [Security Requirements](#332-security-requirements)  
       - [Reliability & Availability](#333-reliability--availability)  
       - [Maintainability & Supportability](#334-maintainability--supportability)  
       - [Usability](#335-usability)  
       - [Portability](#336-portability)  
       - [Scalability](#337-scalability)  
4. [System Models](#4-system-models)  
   1. [Use Case Diagrams](#41-use-case-diagrams)  
   2. [Activity Diagrams](#42-activity-diagrams)  
   3. [Class Diagram](#43-class-diagram)  
5. [Appendices](#5-appendices)  
   1. [Glossary](#51-glossary)  
   2. [Acronyms](#52-acronyms)  
   3. [Additional Diagrams](#53-additional-diagrams)  

---

## 1. Introduction

### 1.1 Purpose
The purpose of this Software Requirements Specification (SRS) document is to provide a detailed, industry-standard description of the requirements for the Indian Sign Language Translator App. This app will translate user-provided speech or text into dynamic 3D hand-model visualizations representing Indian Sign Language (ISL). This document serves as a reference for project stakeholders, designers, developers, and testers.

### 1.2 Scope
The Indian Sign Language (ISL) Translator App is a cross-platform mobile application designed to facilitate communication between hearing individuals and the deaf or hard-of-hearing (HoH) community by translating spoken or typed text into Indian Sign Language through animated 3D hand gestures. The application aims to foster inclusivity and enhance accessibility for ISL users.

This system encompasses the following components:
- **Speech and Text Input Modules**: To capture and preprocess user input.
- **Natural Language Processing (NLP) Engine**: To interpret and convert inputs to ISL gesture codes.
- **3D Animation Module**: To render and animate lifelike hand gestures that represent ISL signs.
- **Mobile Interface**: Intuitive UI for users to interact with the system on Android and iOS platforms.
- **Backend Infrastructure**: Cloud services and APIs supporting translation processing, user analytics, and content updates.

**Key Deliverables**:
- Functional mobile application (Android & iOS)
- NLP-driven translation backend with API interfaces
- Comprehensive 3D gesture asset library
- User and developer documentation including setup guides and API references

**Out-of-Scope Elements**:
- Real-time video-to-sign interpretation
- Multi-language support beyond Indian Sign Language
- Haptic feedback or tactile output features
- Integration with wearable assistive technologies (e.g., smart gloves)

#### Audience

This document is intended for the following stakeholders:

| Stakeholder Type         | Description                                                                 |
|--------------------------|-----------------------------------------------------------------------------|
| **Developers**           | Responsible for implementing application features and maintaining codebase. |
| **Designers**            | Contribute to the UI/UX and gesture animation system.                       |
| **Test Engineers**       | Validate software functionality and non-functional attributes.              |
| **Project Managers**     | Oversee the development lifecycle and track deliverables.                   |
| **Accessibility Advocates** | Provide domain expertise to ensure ISL accuracy and inclusivity.         |
| **End Users**            | Individuals who will utilize the app for real-world communication, including learners and native ISL users. |
| **Regulatory/Compliance Personnel** | Ensure adherence to accessibility standards and data privacy regulations. |
 

### 1.3 Definitions, Acronyms, and Abbreviations
| Term              | Definition                                                                 |
|-------------------|-----------------------------------------------------------------------------|
| ISL               | Indian Sign Language                                                        |
| SRS               | Software Requirements Specification                                          |
| UI                | User Interface                                                              |
| UX                | User Experience                                                             |
| API               | Application Programming Interface                                           |
| NLP               | Natural Language Processing                                                 |
| TTS               | Text-to-Speech (optional for future enhancements)                            |

### 1.4 References
- IEEE Std 830-1998, IEEE Recommended Practice for Software Requirements Specifications.  
- ISL Dictionary and Gesture Reference by Indian Sign Language Research and Training Center.  
- Unity3D Documentation: Animation & Rendering.  
- TensorFlow Lite: On-device Speech Recognition.

### 1.5 Overview
This SRS is organized as follows:
- **Section 1**: Introduction and high-level context.  
- **Section 2**: Overall description of the product.  
- **Section 3**: Detailed external interfaces, functional and non-functional requirements.  
- **Section 4**: System models (use cases, UML diagrams).  
- **Section 5**: Appendices and glossary.  

## 2. Overall Description

### 2.1 Product Perspective
The ISL Translator App is a self-contained mobile solution, integrating speech/text input modules, an NLP translation engine, and 3D rendering. It will interface with:
- Mobile device sensors (microphone, touch screen).  
- Local storage for caching models and assets.  
- Optional cloud API for enhanced translation accuracy.

### 2.2 Product Functions
1. **Input Capture**: Record speech; accept typed text.  
2. **Preprocessing**: Noise reduction, tokenization, normalization.  
3. **Translation Engine**: Map input tokens to ISL grammar and gesture codes.  
4. **3D Rendering**: Load hand-model assets; animate gestures.  
5. **User Settings**: Language selection, rendering quality, offline/online mode.  
6. **Logging & Analytics**: Usage metrics, error reporting.

### 2.3 User Classes and Characteristics
| User Class           | Description                                                                                     |
|----------------------|-------------------------------------------------------------------------------------------------|
| End User             | Hearing individuals seeking to communicate via ISL or learn gestures. Moderate technical skill. |
| Deaf/HoH Advocate    | Users fluent in ISL verifying translation fidelity.                                             |
| Developer            | Software engineers integrating API or extending functionality.                                   |
| Administrator        | Manages cloud resources, analytics, user accounts.                                              |

### 2.4 Operating Environment
- **Mobile OS**: Android 8.0+; iOS 12.0+.
- **Frontend**: Flutter framework for building multi-OS binaries (Android, iOS, Web, Desktop).
- **Backend**: FastAPI (Python) server hosted on AWS or similar.  
- **3D Engine**: Unity3D or Unreal Engine integrated via SDK.  
- **Machine Learning**: TensorFlow Lite for on-device models; TensorFlow Serving for cloud.

### 2.5 Design and Implementation Constraints
- On-device speech recognition limited by mobile CPU and memory.  
- 3D assets must be optimized (<10 MB total) to minimize app size.  
- Compliance with app store policies (Google Play, Apple App Store).

### 2.6 Assumptions and Dependencies
- Users have internet connectivity for cloud-based translations.  
- Device has a functional microphone and sufficient GPU for 3D rendering.  
- ISL gesture library covers primary vocabulary (~2000 signs).

## 3. Specific Requirements

### 3.1 External Interface Requirements

#### 3.1.1 User Interfaces
- **Home Screen**: Options for speech or text input, recent translations.  
- **Input Screen**: Record button with visual feedback; text field with real-time suggestions.  
- **Visualization Screen**: 3D viewport with playback controls (play, pause, rewind).  
- **Settings Screen**: Toggle offline mode, rendering quality slider, language preferences.

#### 3.1.2 Hardware Interfaces
- **Microphone**: Access via OS APIs for PCM audio input.  
- **Touch Screen**: Standard multi-touch gestures for UI navigation.

#### 3.1.3 Software Interfaces
- **Speech-to-Text SDK**: TensorFlow Lite or platform-native APIs.  
- **Translation API**: REST endpoints for cloud translation (e.g., `/translate/text`, `/translate/speech`).  
- **3D Rendering SDK**: Unity3D plugin or native OpenGL ES interface.

#### 3.1.4 Communication Interfaces
- **HTTPS**: Secure RESTful API calls over TLS 1.2+.  
- **WebSockets (Optional)**: Real-time streaming for long session translations.

### 3.2 Functional Requirements
| ID    | Requirement                                                                                               |
|-------|-----------------------------------------------------------------------------------------------------------|
| FR-1  | The system shall capture speech input and convert it to text with ≥90% accuracy in quiet environments.     |
| FR-2  | The system shall accept text input up to 500 characters.                                                  |
| FR-3  | The system shall translate recognized text into a sequence of ISL gesture codes within 2 seconds.         |
| FR-4  | The system shall render each gesture in 3D at 60 FPS on supported devices.                                |
| FR-5  | The system shall allow users to replay, pause, or rewind gesture animations.                              |
| FR-6  | The system shall cache the last 50 translations for offline access.                                       |
| FR-7  | The system shall log translation errors and send anonymized analytics to the backend server.              |

### 3.3 Non-functional Requirements

#### 3.3.1 Performance Requirements
- End-to-end translation latency ≤3 seconds (speech to rendered animation).  
- App startup time ≤2 seconds on modern devices.

#### 3.3.2 Security Requirements
- All stored data (cached translations, user settings) shall be encrypted at rest.  
- All network communications shall use HTTPS/TLS 1.2+.  
- User analytics data shall be anonymized.

#### 3.3.3 Reliability & Availability
- System uptime ≥99.5% for cloud services.  
- On-device mode shall function without network connectivity.

#### 3.3.4 Maintainability & Supportability
- Modular codebase with clear separation between input processing, translation engine, and rendering.  
- Automated unit and integration tests covering ≥80% of code.

#### 3.3.5 Usability
- UI shall conform to Material Design (Android) and Human Interface Guidelines (iOS).  
- Gesture visualizations shall include optional captions and descriptions.  
- Onboarding tutorial explaining core features.

#### 3.3.6 Portability
- Codebase shall be adaptable to emerging platforms (e.g., UWP, web) with ≤20% rework.

#### 3.3.7 Scalability
- Backend translation service shall support up to 10,000 concurrent users with autoscaling.

## 4. System Models

### 4.1 Use Case Diagrams
![Use Case Diagram](./diagrams/use_case.png)

### 4.2 Activity Diagrams
![Activity Diagram](./diagrams/activity_diagram.png)

### 4.3 Class Diagram
![Class Diagram](./diagrams/class_diagram.png)

> *Note: Detailed UML diagrams are provided in the appendices.*

## 5. Appendices

### 5.1 Glossary
| Term          | Definition                                                |
|---------------|-----------------------------------------------------------|
| Gesture Code  | An internal representation mapping a concept to an animation. |
| Renderer      | Software module for drawing 3D hand-models.              |

### 5.2 Acronyms
- ISL: Indian Sign Language  
- API: Application Programming Interface  
- UI: User Interface  

### 5.3 Additional Diagrams
- Sequence Diagrams (Appendix A)  
- Data Flow Diagrams (Appendix B)  
- State Machine Diagrams (Appendix C)  

---

*End of SRS Document*
